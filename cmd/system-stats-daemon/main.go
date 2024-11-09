package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.GetInt("n"))
}
