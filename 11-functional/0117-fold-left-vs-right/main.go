package main

import "fmt"

func foldl(f func(int, int) int, init int, xs []int) int {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func foldr(f func(int, int) int, init int, xs []int) int {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
	}
	return acc
}

func main() {
	sub := func(a, b int) int { return a - b }
	xs := []int{1, 2, 3}
	fmt.Println(foldl(sub, 0, xs), foldr(sub, 0, xs))
}
