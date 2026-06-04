package main

import "fmt"

func glob(pattern, text string) bool {
	var match func(p, t int) bool
	match = func(p, t int) bool {
		if p == len(pattern) {
			return t == len(text)
		}
		if pattern[p] == '*' {
			for k := t; k <= len(text); k++ {
				if match(p+1, k) {
					return true
				}
			}
			return false
		}
		if t < len(text) && pattern[p] == text[t] {
			return match(p+1, t+1)
		}
		return false
	}
	return match(0, 0)
}

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(label(glob("a*b", "aaab")), label(glob("a*b", "aac")))
}
