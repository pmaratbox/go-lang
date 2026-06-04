package main

import "fmt"

func minMax(xs []int) (int, int) {
	lo, hi := xs[0], xs[0]
	for _, x := range xs[1:] {
		if x < lo {
			lo = x
		}
		if x > hi {
			hi = x
		}
	}
	return lo, hi
}

func main() {
	lo, hi := minMax([]int{4, 1, 7})
	fmt.Println(lo, hi)
}
