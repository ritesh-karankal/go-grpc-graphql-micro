package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/catalog"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/telemetry"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
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

	shutdown, err := telemetry.Init(ctx, "catalog-service")
	if err != nil {
		log.Fatal(err)
	}
	defer telemetry.Flush(shutdown)

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()

	if err := r.Setup(ctx); err != nil {
		log.Fatal(err)
	}

	s := catalog.NewService(r)
	if err := catalog.SeedSampleProducts(ctx, s); err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 8080...")
	if err := catalog.ListenGRPC(ctx, s, 8080); err != nil {
		log.Println(err)
	}
	log.Println("Stopped")
}
