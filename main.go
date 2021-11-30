package main

import (
	"context"

	"github.com/danial2026/file-sharing-go/app"

	kafkaApp "github.com/danial2026/file-sharing-go/controllers/kafka"
)

func main() {
	ctx := context.Background()

	go kafkaApp.Consume(ctx)
	app.StartApplication()
}
