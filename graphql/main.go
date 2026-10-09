package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Cancelled on SIGTERM: the server then drains, and telemetry is flushed last
	ctx, stop := lifecycle.SignalContext()
	defer stop()

	shutdown, err := telemetry.Init(ctx, "graphql-gateway")
	if err != nil {
		log.Fatal(err)
	}
	defer flush(shutdown)

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	}

	graphqlServer := handler.New(s.ToExecutableSchema())

	graphqlServer.AddTransport(transport.Options{})
	graphqlServer.AddTransport(transport.GET{})
	graphqlServer.AddTransport(transport.POST{})

	// One span per GraphQL operation and per resolver call. Variables are left out
	// so user input (names etc.) doesn't end up in the traces.
	graphqlServer.Use(otelgqlgen.Middleware(
		otelgqlgen.WithoutVariables(),
		otelgqlgen.WithCreateSpanFromFields(func(fc *graphql.FieldContext) bool {
			return fc.IsResolver
		}),
	))

	// Per-operation count and latency with an ok / client_error / server_error outcome:
	// the SLIs. HTTP metrics can't provide them, since GraphQL errors still return 200.
	opMetrics, err := newOperationMetrics(otel.GetMeterProvider())
	if err != nil {
		log.Fatal(err)
	}
	graphqlServer.Use(opMetrics)

	mux := http.NewServeMux()
	mux.Handle("/graphql", otelhttp.NewHandler(corsMiddleware(syntheticMiddleware(graphqlServer)), "graphql"))
	mux.Handle("/playground", corsMiddleware(playground.Handler("ritesh", "/graphql")))

	srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 10 * time.Second}
	if err := lifecycle.ServeHTTP(ctx, srv); err != nil {
		log.Println(err)
	}
	log.Println("Stopped")
}

// flush sends buffered spans, metrics and logs before the process exits.
func flush(shutdown func(context.Context) error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shutdown(ctx); err != nil {
		log.Println("Failed to flush telemetry:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
