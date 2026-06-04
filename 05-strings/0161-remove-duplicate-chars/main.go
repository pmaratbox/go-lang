package main

import (
	"fmt"
	"strings"
)

func main() {
	seen := map[rune]bool{}
	var b strings.Builder
	for _, r := range "aabbcc" {
		if !seen[r] {
			seen[r] = true
			b.WriteRune(r)
		}
	}
	fmt.Println(b.String())
}
