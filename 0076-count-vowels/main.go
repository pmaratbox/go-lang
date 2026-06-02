package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "hello"
	vowels := "aeiou"
	count := 0
	for _, ch := range word {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	fmt.Println(count)
}
