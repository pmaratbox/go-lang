package main

import "fmt"

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

func main() {
	fmt.Println(modPow(3, 13, 7))
}
