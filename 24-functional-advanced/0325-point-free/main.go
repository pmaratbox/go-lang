package main

import "fmt"

func mapInts(f func(int) int) func([]int) []int {
	return func(xs []int) []int {
		out := make([]int, len(xs))
		for i, x := range xs {
			out[i] = f(x)
		}
		return out
	}
}

func sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

// compose builds (g . f).
func compose(g func([]int) int, f func([]int) []int) func([]int) int {
	return func(xs []int) int { return g(f(xs)) }
}

func main() {
	square := func(x int) int { return x * x }
	sumOfSquares := compose(sum, mapInts(square))
	fmt.Println(sumOfSquares([]int{1, 2, 3}))
}
