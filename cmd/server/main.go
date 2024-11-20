package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	application, err := app.NewServer(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Starting gRPC server")
		err := application.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-sigC
	log.Println("Stopping gRPC server")
	cancel()
	application.Stop()
}
