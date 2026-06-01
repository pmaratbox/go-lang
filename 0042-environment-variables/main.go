package main

import (
	"fmt"
	"os"
)

func main() {
	value, ok := os.LookupEnv("LESSON_ENV_VAR")
	if !ok {
		value = "default"
	}
	fmt.Printf("value: %s\n", value)
}
