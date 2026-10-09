package main

import (
	"context"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func operationCtx(opType ast.Operation, fields ...string) context.Context {
	op := &ast.OperationDefinition{Operation: opType}
	for _, f := range fields {
		op.SelectionSet = append(op.SelectionSet, &ast.Field{Name: f, Alias: f})
	}
	return graphql.WithOperationContext(context.Background(), &graphql.OperationContext{
		Operation: op,
		Stats:     graphql.Stats{OperationStart: time.Now()},
	})
}

func resolverErr(err error) *gqlerror.Error {
	return &gqlerror.Error{Message: err.Error(), Err: err}
}

func TestOperationMetricsOutcome(t *testing.T) {
	tests := []struct {
		name          string
		ctx           context.Context
		errs          gqlerror.List
		wantOperation string
		wantType      string
		wantOutcome   string
	}{
		{"success", operationCtx(ast.Mutation, "createOrder"), nil,
			"createOrder", "mutation", outcomeOK},
		{"unknown account is the caller's mistake", operationCtx(ast.Mutation, "createOrder"),
			gqlerror.List{resolverErr(status.Error(codes.NotFound, "account not found"))},
			"createOrder", "mutation", outcomeClientError},
		{"bad quantity rejected by the gateway", operationCtx(ast.Mutation, "createOrder"),
			gqlerror.List{resolverErr(ErrInvalidParameter)},
			"createOrder", "mutation", outcomeClientError},
		{"dependency down is ours", operationCtx(ast.Mutation, "createOrder"),
			gqlerror.List{resolverErr(status.Error(codes.Unavailable, "could not load products"))},
			"createOrder", "mutation", outcomeServerError},
		{"plain error from a service is ours", operationCtx(ast.Query, "products"),
			gqlerror.List{resolverErr(status.Error(codes.Unknown, "boom"))},
			"products", "query", outcomeServerError},
		{"one server error makes the whole operation a server error", operationCtx(ast.Query, "accounts"),
			gqlerror.List{
				resolverErr(status.Error(codes.NotFound, "account not found")),
				resolverErr(status.Error(codes.Internal, "could not get orders")),
			},
			"accounts", "query", outcomeServerError},
		{"several root fields, sorted", operationCtx(ast.Query, "products", "accounts"), nil,
			"accounts,products", "query", outcomeOK},
		{"introspection", operationCtx(ast.Query, "__schema"), nil,
			"introspection", "query", outcomeOK},
		{"parse or validation failure has no operation context", context.Background(),
			gqlerror.List{{Message: "Cannot query field \"nope\" on type \"Query\"."}},
			"invalid", "unknown", outcomeClientError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := sdkmetric.NewManualReader()
			m, err := newOperationMetrics(sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)))
			if err != nil {
				t.Fatal(err)
			}

			m.InterceptResponse(tt.ctx, func(context.Context) *graphql.Response {
				return &graphql.Response{Errors: tt.errs}
			})

			var rm metricdata.ResourceMetrics
			if err := reader.Collect(context.Background(), &rm); err != nil {
				t.Fatal(err)
			}
			want := attribute.NewSet(
				attribute.String("operation", tt.wantOperation),
				attribute.String("operation_type", tt.wantType),
				attribute.String("outcome", tt.wantOutcome),
			)
			assertCounted(t, rm, want)
		})
	}
}

// assertCounted checks that both instruments recorded exactly one operation with the
// expected attributes.
func assertCounted(t *testing.T, rm metricdata.ResourceMetrics, want attribute.Set) {
	t.Helper()
	found := map[string]bool{}
	for _, sm := range rm.ScopeMetrics {
		for _, md := range sm.Metrics {
			switch data := md.Data.(type) {
			case metricdata.Sum[int64]:
				if len(data.DataPoints) != 1 || data.DataPoints[0].Value != 1 {
					t.Fatalf("%s: want 1 data point with value 1, got %+v", md.Name, data.DataPoints)
				}
				if got := data.DataPoints[0].Attributes; !got.Equals(&want) {
					t.Fatalf("%s: attributes = %v, want %v", md.Name, got.ToSlice(), want.ToSlice())
				}
			case metricdata.Histogram[float64]:
				if len(data.DataPoints) != 1 || data.DataPoints[0].Count != 1 {
					t.Fatalf("%s: want 1 data point with count 1, got %+v", md.Name, data.DataPoints)
				}
				if got := data.DataPoints[0].Attributes; !got.Equals(&want) {
					t.Fatalf("%s: attributes = %v, want %v", md.Name, got.ToSlice(), want.ToSlice())
				}
			}
			found[md.Name] = true
		}
	}
	for _, name := range []string{"graphql.server.operations", "graphql.server.operation.duration"} {
		if !found[name] {
			t.Errorf("metric %s not recorded", name)
		}
	}
}
