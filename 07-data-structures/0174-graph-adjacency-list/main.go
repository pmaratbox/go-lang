package main

import (
	"fmt"
	"strings"
)

func main() {
	adj := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},
	}
	var parts []string
	for _, n := range adj[0] {
		parts = append(parts, fmt.Sprint(n))
	}
	fmt.Println(strings.Join(parts, " "))
}
