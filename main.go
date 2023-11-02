package main

import (
	"context"
	"log"

	"github.com/go-grpc-http/server"
	"github.com/rohanraj7316/ds/migrations"
	"github.com/rohanraj7316/ds/resources/stocks"
	"github.com/rohanraj7316/ds/utils/postgres"
	"github.com/rohanraj7316/logger"
)

func main() {
	err := logger.Configure()
	if err != nil {
		log.Fatalf("failed to initialize logger: %s", err)
	}

	ps, err := postgres.New()
	if err != nil {
		log.Fatalf("failed to initialize postgres: %s", err)
	}

	err = migrations.Migration(context.Background(), ps)
	if err != nil {
		log.Fatalf("failed to initialize postgres: %s", err)
	}

	sH, err := stocks.New(ps)
	if err != nil {
		log.Fatalf("failed to initialize stock handler: %s", err)
	}

	srv, err := server.New(
		server.WithDualRegisterer(sH),
		server.WithReflection(true),
	)
	if err != nil {
		log.Fatalf("failed to initialize server: %s", err)
	}

	err = srv.Run(context.Background())
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
