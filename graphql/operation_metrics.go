package main

import (
	"context"
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Outcomes of a GraphQL operation, for SLIs. gqlgen answers HTTP 200 even when a
// resolver fails, so HTTP status metrics can't tell a failed checkout from a good one.
const (
	outcomeOK          = "ok"
	outcomeClientError = "client_error" // the caller's mistake: doesn't count against the SLO
	outcomeServerError = "server_error" // ours: burns the error budget
)

// operationMetrics records, per GraphQL operation:
//
//	graphql.server.operations          counter    {operation, operation_type, outcome, synthetic}
//	graphql.server.operation.duration  histogram  {operation, operation_type, outcome, synthetic}
//
// In Prometheus: graphql_server_operations_total and
// graphql_server_operation_duration_seconds_bucket.
//
// The same attributes go on the GraphQL operation span (gqlgen runs extensions in
// registration order from the outside in, so this one runs inside otelgqlgen's span),
// so an SLO panel and the traces behind it can be filtered with the same labels.
type operationMetrics struct {
	count    metric.Int64Counter
	duration metric.Float64Histogram
}

var (
	_ graphql.HandlerExtension    = (*operationMetrics)(nil)
	_ graphql.ResponseInterceptor = (*operationMetrics)(nil)
)

func newOperationMetrics(mp metric.MeterProvider) (*operationMetrics, error) {
	meter := mp.Meter("github.com/ritesh-karankal/go-grpc-graphql-micro/graphql")

	count, err := meter.Int64Counter("graphql.server.operations",
		metric.WithDescription("GraphQL operations handled, by root field and outcome"),
		metric.WithUnit("{operation}"))
	if err != nil {
		return nil, err
	}

	// Bucket bounds include the SLO latency thresholds (0.5s, 0.8s, 1s), so
	// "share of requests faster than X" can be read exactly from the buckets.
	duration, err := meter.Float64Histogram("graphql.server.operation.duration",
		metric.WithDescription("Time to handle a GraphQL operation"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.05, 0.1, 0.25, 0.5, 0.8, 1, 2.5, 5, 10))
	if err != nil {
		return nil, err
	}

	return &operationMetrics{count: count, duration: duration}, nil
}

func (m *operationMetrics) ExtensionName() string { return "OperationMetrics" }

func (m *operationMetrics) Validate(graphql.ExecutableSchema) error { return nil }

func (m *operationMetrics) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()
	operation, opType := "invalid", "unknown" // parse/validation failures have no operation context
	if graphql.HasOperationContext(ctx) {
		oc := graphql.GetOperationContext(ctx)
		start = oc.Stats.OperationStart
		if oc.Operation != nil {
			operation, opType = rootFields(oc.Operation), string(oc.Operation.Operation)
		}
	}

	resp := next(ctx)

	outcome := outcomeOK
	if resp != nil {
		outcome = classify(resp.Errors)
	}
	if operation == "invalid" && outcome == outcomeOK {
		outcome = outcomeClientError
	}

	kv := []attribute.KeyValue{
		attribute.String("operation", operation),
		attribute.String("operation_type", opType),
		attribute.String("outcome", outcome),
		attribute.Bool("synthetic", isSynthetic(ctx)),
	}
	trace.SpanFromContext(ctx).SetAttributes(kv...)

	attrs := metric.WithAttributes(kv...)
	m.count.Add(ctx, 1, attrs)
	m.duration.Record(ctx, time.Since(start).Seconds(), attrs)
	return resp
}

// rootFields names an operation after its top-level fields ("createOrder",
// "accounts,products"), which the schema bounds, unlike client-chosen operation names.
func rootFields(op *ast.OperationDefinition) string {
	seen := map[string]bool{}
	for _, sel := range op.SelectionSet {
		f, ok := sel.(*ast.Field)
		if !ok {
			continue
		}
		name := f.Name
		if strings.HasPrefix(name, "__") {
			name = "introspection"
		}
		seen[name] = true
	}
	if len(seen) == 0 {
		return "unknown"
	}
	names := make([]string, 0, len(seen))
	for n := range seen {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}

// classify returns server_error if any error is ours, client_error if all errors are
// the caller's, and ok if there are none.
func classify(errs gqlerror.List) string {
	if len(errs) == 0 {
		return outcomeOK
	}
	for _, e := range errs {
		if !isClientError(e) {
			return outcomeServerError
		}
	}
	return outcomeClientError
}

func isClientError(e *gqlerror.Error) bool {
	if e.Err == nil {
		// gqlgen's own errors (bad query, wrong variable types) come without a cause
		return true
	}
	if errors.Is(e.Err, ErrInvalidParameter) {
		return true
	}
	var se interface{ GRPCStatus() *status.Status }
	if !errors.As(e.Err, &se) {
		return false // unexpected error inside the gateway
	}
	switch se.GRPCStatus().Code() {
	case codes.NotFound, codes.InvalidArgument, codes.AlreadyExists, codes.FailedPrecondition,
		codes.OutOfRange, codes.PermissionDenied, codes.Unauthenticated:
		return true
	default:
		return false
	}
}

// Synthetic traffic (load tests, probes) sends "X-Synthetic: true". It is recorded as
// synthetic=true so SLOs can count real users only. Anyone can set the header; it only
// affects how requests are counted, not how they are served.

type syntheticKey struct{}

func syntheticMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.EqualFold(r.Header.Get("X-Synthetic"), "true") {
			r = r.WithContext(context.WithValue(r.Context(), syntheticKey{}, true))
		}
		next.ServeHTTP(w, r)
	})
}

func isSynthetic(ctx context.Context) bool {
	v, _ := ctx.Value(syntheticKey{}).(bool)
	return v
}
