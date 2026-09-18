package main

import (
	"context"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ritesh-karankal/go-grpc-graphql-micro/catalog"
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

	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()

	ctx := context.Background()
	if err := r.Setup(ctx); err != nil {
		log.Fatal(err)
	}

	s := catalog.NewService(r)
	if err := catalog.SeedSampleProducts(ctx, s); err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port 8080...")
	log.Fatal(catalog.ListenGRPC(s, 8080))

}