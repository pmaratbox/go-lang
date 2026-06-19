package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
	fmt.Printf("%s:%d\n", viper.GetString("server.host"), viper.GetInt("server.port"))
}
