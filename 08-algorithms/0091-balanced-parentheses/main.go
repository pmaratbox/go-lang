package main

import "fmt"

func isBalanced(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	for _, s := range []string{"(())", "(()"} {
		if isBalanced(s) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
