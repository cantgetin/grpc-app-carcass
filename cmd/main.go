package main

import (
	"context"
	"grpc-carcass/internal/app"
	"log"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		log.Fatalf("error running grpc server, %v", err)
	}
}
