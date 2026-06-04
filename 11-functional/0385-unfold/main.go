package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := 1
	var out []string
	for i := 0; i < 5; i++ {
		out = append(out, strconv.Itoa(x))
		x *= 2
	}
	fmt.Println(strings.Join(out, " "))
}
