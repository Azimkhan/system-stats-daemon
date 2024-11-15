package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.GetInt("n"))
}
