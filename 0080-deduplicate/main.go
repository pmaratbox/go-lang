package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2, 2, 3, 1}
	seen := map[int]bool{}
	unique := []int{}
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			unique = append(unique, n)
		}
	}
	parts := []string{}
	for _, x := range unique {
		parts = append(parts, fmt.Sprint(x))
	}
	fmt.Println(strings.Join(parts, " "))
}
