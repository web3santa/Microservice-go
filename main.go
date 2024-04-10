package main

import (
	"context"
	"log"
	"microservice/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		log.Panic(err)
	}

}
