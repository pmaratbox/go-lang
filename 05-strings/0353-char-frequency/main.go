package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aab"
	counts := map[rune]int{}
	var order []rune
	for _, c := range s {
		if _, ok := counts[c]; !ok {
			order = append(order, c)
		}
		counts[c]++
	}
	var parts []string
	for _, c := range order {
		parts = append(parts, fmt.Sprintf("%c:%d", c, counts[c]))
	}
	fmt.Println(strings.Join(parts, " "))
}
