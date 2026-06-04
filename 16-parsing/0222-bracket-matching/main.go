package main

import "fmt"

func balanced(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func say(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(say(balanced("([{}])")), say(balanced("([)]")))
}
