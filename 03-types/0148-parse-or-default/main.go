package main

import (
	"fmt"
	"strconv"
)

// parseOrDefault returns the parsed int, or the default on failure.
func parseOrDefault(s string, def int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return def
}

func main() {
	a := parseOrDefault("42", 0)
	b := parseOrDefault("x", 0)

	fmt.Printf("%d %d\n", a, b)
}
