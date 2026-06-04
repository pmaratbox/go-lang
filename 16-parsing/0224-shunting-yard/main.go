package main

import (
	"fmt"
	"strings"
)

func prec(op string) int {
	if op == "*" || op == "/" {
		return 2
	}
	return 1
}

func main() {
	expr := "3 + 4 * 2"
	var out []string
	var ops []string
	for _, tok := range strings.Fields(expr) {
		switch tok {
		case "+", "-", "*", "/":
			for len(ops) > 0 && prec(ops[len(ops)-1]) >= prec(tok) {
				out = append(out, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, tok)
		default:
			out = append(out, tok)
		}
	}
	for len(ops) > 0 {
		out = append(out, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	fmt.Println(strings.Join(out, " "))
}
