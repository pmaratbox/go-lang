package main

import "fmt"

func fastPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func main() {
	fmt.Println(fastPow(2, 10))
}
