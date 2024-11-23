package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/Azimkhan/system-stats-daemon/internal/config"
	"github.com/Azimkhan/system-stats-daemon/internal/logging"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("addr", ":50051")
	viper.SetDefault("connTimeout", 5*time.Second)
}

func main() {
	err := viper.ReadInConfig()
	if err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			fmt.Println("Config file not found, using defaults")
		} else {
			fmt.Printf("Error reading config, %v\n", err)
			return
		}
	}

	var addr string
	err = viper.UnmarshalKey("addr", &addr)
	if err != nil {
		fmt.Printf("Error reading addr, %v\n", err)
		return
	}

	var connTimeout time.Duration
	err = viper.UnmarshalKey("connTimeout", &connTimeout)
	if err != nil {
		fmt.Printf("Error reading connTimeout, %v\n", err)
		return
	}

	log.Printf("Connecting to %s\n", addr)
	logger, err := logging.NewLogger(&config.LoggingConfig{
		Level:  "info",
		Format: "text",
	})
	if err != nil {
		fmt.Printf("Error creating logger, %v\n", err)
		return
	}
	application, err := app.NewClientApp(addr, connTimeout, logger)
	if err != nil {
		fmt.Printf("Error creating application, %v\n", err)
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
