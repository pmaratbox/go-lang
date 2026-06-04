package main

import (
	"fmt"
	"strings"
)

func main() {
	values := []int{3, 1, 2}
	for _, n := range values {
		fmt.Println(strings.Repeat("#", n))
	}
}
