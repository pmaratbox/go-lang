package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	os.Setenv("NAME", "from-env")

	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
	viper.BindEnv("name", "NAME")

	fmt.Println(viper.GetString("name"))
}
