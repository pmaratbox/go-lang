package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word := "banana"
	counts := map[rune]int{}
	for _, ch := range word {
		counts[ch]++
	}

	keys := []rune{}
	for ch := range counts {
		keys = append(keys, ch)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	parts := []string{}
	for _, ch := range keys {
		parts = append(parts, fmt.Sprintf("%c:%d", ch, counts[ch]))
	}
	fmt.Println(strings.Join(parts, " "))
}
