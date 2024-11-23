package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Printf("failed to read config: %v\n", err)
		os.Exit(1)
	}

	logger, err := logging.NewLogger(conf.Logging)
	if err != nil {
		fmt.Printf("failed to create logger: %v\n", err)
		os.Exit(1)
	}

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	application, err := app.NewServer(ctx, conf.Server, conf.Stream, logger)
	if err != nil {
		logger.Error("failed to create application", "error", err)
	}

	go func() {
		err := application.Serve()
		if err != nil {
			logger.Error("failed to serve", "error", err)
		}
	}()

	<-sigC
	logger.Info("shutting down")
	cancel()
	application.Stop()
}
