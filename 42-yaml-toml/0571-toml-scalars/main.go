package main

import (
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	var t map[string]any
	if err := toml.Unmarshal([]byte("title = \"demo\"\nversion = 2\n"), &t); err != nil {
		panic(err)
	}
	// title is a string, version is an integer (int64) — both print plainly.
	fmt.Printf("%v %v\n", t["title"], t["version"])
}
