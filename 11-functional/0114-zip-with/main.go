package main

import (
	"fmt"
	"strings"
)

func zipWith(f func(int, int) int, a, b []int) []int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = f(a[i], b[i])
	}
	return out
}

func main() {
	r := zipWith(func(x, y int) int { return x + y }, []int{1, 2, 3}, []int{4, 5, 6})
	parts := make([]string, len(r))
	for i, v := range r {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
