package main

import (
	"context"
	"log"
	"microservice/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Panic(err)
	}

	cancel()

}
