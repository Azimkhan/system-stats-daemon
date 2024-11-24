package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
)

func main() {
	var addr string
	var connTimeout time.Duration

	flag.StringVar(&addr, "addr", ":50051", "server address")
	flag.DurationVar(&connTimeout, "connTimeout", 5*time.Second, "connection timeout")
	flag.Parse()

	log.Printf("Connecting to %s\n", addr)
	logger, err := logging.NewLogger(&config.LoggingConfig{
		Level:   "info",
		Handler: "text",
	})
	if err != nil {
		fmt.Printf("Error creating logger, %v\n", err)
		return
	}
	application, err := app.NewClientApp(addr, connTimeout, logger)
	if err != nil {
		logger.Error("Error creating application", "error", err)
		return
	}
	defer func() {
		err := application.Close()
		if err != nil {
			log.Println(err)
			return
		}

		logger.Info("Application finished.")
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	appErr := make(chan error)
	go func() {
		appErr <- application.Run(ctx)
	}()

	select {
	case <-sigC:
		logger.Info("Interrupted by signal")
	case err := <-appErr:
		if err != nil {
			logger.Error("Error running application, %v\n", "error", err)
		}
	}
	cancel()
}
