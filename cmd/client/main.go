package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Azimkhan/system-stats-daemon/internal/app"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("addr", ":50051")
	viper.SetDefault("connTimeout", 5*time.Second)
}

func main() {
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config, %v\n", err)
		return
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
	application, err := app.NewClientApp(addr, connTimeout)
	defer func() {
		err := application.Close()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Application finished.")
	}()
	if err != nil {
		log.Println(err)
		return
	}
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
		log.Println("Interrupted by signal")
	case err := <-appErr:
		if err != nil {
			fmt.Printf("Error running application, %v\n", err)
		}
	}
	cancel()
}
