package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Split("hello world", " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}
