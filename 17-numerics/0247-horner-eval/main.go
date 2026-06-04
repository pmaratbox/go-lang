package main

import "fmt"

func horner(coeffs []int, x int) int {
	result := 0
	for _, c := range coeffs {
		result = result*x + c
	}
	return result
}

func main() {
	// 2x^2 + 3x + 1 at x=2
	fmt.Println(horner([]int{2, 3, 1}, 2))
}
