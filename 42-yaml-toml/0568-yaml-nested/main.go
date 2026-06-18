package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func main() {
	doc := "server:\n  host: localhost\n  port: 8080\n"

	var d map[string]any
	yaml.Unmarshal([]byte(doc), &d)

	server := d["server"].(map[string]any)
	fmt.Printf("%v:%v\n", server["host"], server["port"])
}
