package main

import "fmt"

func main() {
	pattern := "ab"
	text := "aab"
	state := 0
	match := -1
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[state] {
			state++
		} else if text[i] == pattern[0] {
			state = 1
		} else {
			state = 0
		}
		if state == len(pattern) {
			match = i - len(pattern) + 1
			break
		}
	}
	fmt.Println(match)
}
