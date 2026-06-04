package main

import (
	"fmt"
	"strings"
)

func main() {
	state := "locked"
	events := []string{"coin", "push", "push"}
	var states []string
	for _, e := range events {
		switch {
		case state == "locked" && e == "coin":
			state = "unlocked"
		case state == "unlocked" && e == "push":
			state = "locked"
		case state == "locked" && e == "push":
			state = "locked"
		}
		states = append(states, state)
	}
	fmt.Println(strings.Join(states, " "))
}
