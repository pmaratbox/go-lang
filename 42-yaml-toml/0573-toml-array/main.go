package main

import (
	"fmt"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	var t map[string]any
	if err := toml.Unmarshal([]byte("tags = [\"red\", \"green\", \"blue\"]\n"), &t); err != nil {
		panic(err)
	}
	// A TOML array decodes into a []any; each element here is a string.
	arr := t["tags"].([]any)
	parts := make([]string, len(arr))
	for i, v := range arr {
		parts[i] = v.(string)
	}
	fmt.Println(strings.Join(parts, ","))
}
