package main

import (
	"fmt"
	"strings"
)

func main() {
	tmpl := "hi {name}"
	vars := map[string]string{"name": "Ada"}
	out := tmpl
	for k, v := range vars {
		out = strings.ReplaceAll(out, "{"+k+"}", v)
	}
	fmt.Println(out)
}
