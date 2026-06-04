package main

import (
	"fmt"
	"strings"
)

func main() {
	next := map[string]string{"A": "B", "B": "C", "C": "A"}
	state := "A"
	var visited []string
	for i := 0; i < 3; i++ {
		state = next[state]
		visited = append(visited, state)
	}
	fmt.Println(strings.Join(visited, " "))
}
