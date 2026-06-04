package main

import "fmt"

// fold reduces xs from an identity using combine.
func fold[T any](xs []T, identity T, combine func(T, T) T) T {
	acc := identity
	for _, x := range xs {
		acc = combine(acc, x)
	}
	return acc
}

func main() {
	s := fold([]string{"a", "b", "c"}, "", func(a, b string) string { return a + b })
	n := fold([]int{1, 2, 3}, 0, func(a, b int) int { return a + b })
	fmt.Println(s, n)
}
