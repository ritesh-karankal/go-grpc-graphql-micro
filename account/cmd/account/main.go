package main

import (
	"context"
	"log"
	"time"

	"github.com/ritesh-karankal/go-grpc-graphql-micro/account"
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

	shutdown, err := telemetry.Init(context.Background(), "account-service")
	if err != nil {
		log.Fatal(err)
	}
	telemetry.ShutdownOnSignal(shutdown)

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
	log.Fatal(account.ListenGRPC(s, 8080))
}