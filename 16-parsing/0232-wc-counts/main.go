package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "a b\nc"
	words := len(strings.Fields(text))
	lines := len(strings.Split(text, "\n"))
	chars := len(text)
	fmt.Println(words, lines, chars)
}
