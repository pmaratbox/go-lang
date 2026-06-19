package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
	fmt.Println(strings.Join(viper.GetStringSlice("hosts"), ","))
}
