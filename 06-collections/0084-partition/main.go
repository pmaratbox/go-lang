package main

import (
	"fmt"
	"strings"
)

func join(xs []int) string {
	parts := []string{}
	for _, x := range xs {
		parts = append(parts, fmt.Sprint(x))
	}
	return strings.Join(parts, " ")
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := []int{}
	odds := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		} else {
			odds = append(odds, n)
		}
	}
	fmt.Println("evens: " + join(evens))
	fmt.Println("odds: " + join(odds))
}
