package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("n", 5)
	viper.SetDefault("m", 15)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/system-stats/")
}

func main() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.GetInt("n"))
}
