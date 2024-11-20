package config

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	BindAddr string `mapstructure:"bind-addr"`
}

type StreamingConfig struct {
	InitialDelay time.Duration `mapstructure:"initial-delay"`
	Interval     time.Duration
}
type Config struct {
	Server *ServerConfig
	Stream *StreamingConfig
}

func init() {
	viper.SetDefault("stream.interval", 5*time.Second)
	viper.SetDefault("stream.initialDelay", 15*time.Second)
	viper.SetDefault("server.bindAddr", ":50051")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/system-stats/")
}

func Read() (*Config, error) {
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			log.Fatal(err)
		}
		log.Println("Config file not found, using defaults")
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return conf, nil
}
