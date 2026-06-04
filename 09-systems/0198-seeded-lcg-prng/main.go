package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 1
	var out []string
	for i := 0; i < 3; i++ {
		x = (5*x + 3) % 16
		out = append(out, fmt.Sprintf("%d", x))
	}
	fmt.Println(strings.Join(out, " "))
}
