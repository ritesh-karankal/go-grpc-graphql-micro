package main

import (
	"log"
	"time"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/telemetry"
	"github.com/kelseyhightower/envconfig"
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

	shutdown, err := telemetry.Init(ctx, "account-service")
	if err != nil {
		log.Fatal(err)
	}
	defer telemetry.Flush(shutdown)

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	if err := account.ListenGRPC(ctx, s, 8080); err != nil {
		log.Println(err)
	}
	log.Println("Stopped")
}
