package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
	viper.SetDefault("missing", "fallback")
	fmt.Println(viper.GetString("missing"))
}
