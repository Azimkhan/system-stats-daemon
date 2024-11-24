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
	appErrChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	application, err := app.NewServerApp(ctx, conf.Stats, conf.Server, conf.Stream, logger)
	if err != nil {
		logger.Error("failed to create application", "error", err)
	}
	defer func() {
		cancel()
		application.Stop()
		logger.Info("application stopped")
	}()

	logger.Info(
		"application initialized",
		"stats to collect", conf.Stats,
		"stream initial delay",
		conf.Stream.InitialDelay,
		"stream interval",
		conf.Stream.Interval,
	)

	go func() {
		appErrChan <- application.Serve()
	}()
	select {
	case <-sigC:
		logger.Info("Shut down signal received")
		cancel()

	case appErr := <-appErrChan:
		if appErr != nil {
			logger.Error("application serve error", "error", err)
		}
	}
}
