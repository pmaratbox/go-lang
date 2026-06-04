package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"1", "2", "3"}
	b := []string{"a", "b"}
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	parts := make([]string, n)
	for i := 0; i < n; i++ {
		x, y := "-", "-"
		if i < len(a) {
			x = a[i]
		}
		if i < len(b) {
			y = b[i]
		}
		parts[i] = x + y
	}
	fmt.Println(strings.Join(parts, " "))
}
