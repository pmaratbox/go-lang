package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := 1
	parts := make([]string, 0, 5)
	for n := 0; n < 5; n++ {
		parts = append(parts, strconv.Itoa(c))
		c = c * 2 * (2*n + 1) / (n + 2)
	}
	fmt.Println(strings.Join(parts, " "))
}
