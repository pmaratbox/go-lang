package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world"
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	fmt.Println(strings.Join(words, " "))
}
