package main

import "fmt"

func total(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Printf("sum: %d\n", total(1, 2, 3))
}
