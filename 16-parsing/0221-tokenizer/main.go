package main

import (
	"fmt"
	"strings"
)

func main() {
	expr := "1 + 2"
	var toks []string
	for _, ch := range expr {
		switch {
		case ch >= '0' && ch <= '9':
			toks = append(toks, "NUM")
		case ch == '+':
			toks = append(toks, "PLUS")
		}
	}
	fmt.Println(strings.Join(toks, " "))
}
