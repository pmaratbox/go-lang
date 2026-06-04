package main

import (
	"fmt"
	"strings"
)

func flatMap(xs []int, f func(int) []int) []int {
	out := []int{}
	for _, x := range xs {
		out = append(out, f(x)...)
	}
	return out
}

func main() {
	result := flatMap([]int{1, 2, 3}, func(x int) []int { return []int{x, x * 10} })
	parts := make([]string, len(result))
	for i, v := range result {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
