package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	expr := "3 4 + 5 *"
	var stack []int
	for _, tok := range strings.Fields(expr) {
		switch tok {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var r int
			switch tok {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				r = a / b
			}
			stack = append(stack, r)
		default:
			n, _ := strconv.Atoi(tok)
			stack = append(stack, n)
		}
	}
	fmt.Println(stack[0])
}
