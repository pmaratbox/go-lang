package main

import (
	"fmt"
	"strings"
)

func main() {
	nested := [][]int{{1, 2}, {3, 4}}
	flat := []int{}
	for _, row := range nested {
		flat = append(flat, row...)
	}

	parts := []string{}
	for _, x := range flat {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
