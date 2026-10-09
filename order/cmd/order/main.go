package main

import (
	"context"
	"log"
	"time"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/order"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/telemetry"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"CATALOG_SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Cancelled on SIGTERM: the server then drains, and telemetry is flushed last
	ctx, stop := lifecycle.SignalContext()
	defer stop()

	shutdown, err := telemetry.Init(ctx, "order-service")
	if err != nil {
		log.Fatal(err)
	}
	defer flush(shutdown)

	var r order.Repository

	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()

	log.Println("Listening on port 8080...")
	s := order.NewService(r)
	if err := order.ListenGRPC(ctx, s, cfg.AccountURL, cfg.CatalogURL, 8080); err != nil {
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
