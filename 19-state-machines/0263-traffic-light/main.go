package main

import (
	"fmt"
	"strings"
)

func main() {
	next := map[string]string{"red": "green", "green": "yellow", "yellow": "red"}
	state := "red"
	var states []string
	for i := 0; i < 4; i++ {
		state = next[state]
		states = append(states, state)
	}
	fmt.Println(strings.Join(states, " "))
}
