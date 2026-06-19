package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
	fmt.Println(viper.GetInt("retries"))
}
