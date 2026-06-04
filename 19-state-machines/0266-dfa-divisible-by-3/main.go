package main

import "fmt"

func divisibleBy3(bits string) bool {
	state := 0
	for _, c := range bits {
		b := int(c - '0')
		state = (state*2 + b) % 3
	}
	return state == 0
}

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(label(divisibleBy3("110")), label(divisibleBy3("100")))
}
