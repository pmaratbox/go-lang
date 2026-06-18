package main

import (
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	var t map[string]any
	toml.Unmarshal([]byte("[server]\nhost = \"localhost\"\nport = 8080\n"), &t)
	server := t["server"].(map[string]any)
	fmt.Printf("host=%v port=%v\n", server["host"], server["port"])
}
