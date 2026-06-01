package main

import (
	"fmt"
	"strconv"
	"strings"
)

func squares() func() int {
	n := 0
	return func() int {
		n++
		return n * n
	}
}

func main() {
	next := squares()
	parts := make([]string, 3)
	for i := 0; i < 3; i++ {
		parts[i] = strconv.Itoa(next())
	}
	fmt.Println(strings.Join(parts, " "))
}
