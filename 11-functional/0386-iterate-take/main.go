package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := func(x int) int { return x * 3 }
	x := 1
	var out []string
	for i := 0; i < 4; i++ {
		out = append(out, strconv.Itoa(x))
		x = f(x)
	}
	fmt.Println(strings.Join(out, " "))
}
