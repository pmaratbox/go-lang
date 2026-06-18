package main

import (
	"fmt"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

func main() {
	doc := "[[servers]]\nname = \"alpha\"\n[[servers]]\nname = \"beta\"\n"

	var data map[string]any
	if err := toml.Unmarshal([]byte(doc), &data); err != nil {
		panic(err)
	}

	servers := data["servers"].([]any)
	names := make([]string, 0, len(servers))
	for _, s := range servers {
		names = append(names, s.(map[string]any)["name"].(string))
	}

	fmt.Println(strings.Join(names, ","))
}
