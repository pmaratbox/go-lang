package main

import "fmt"

// factCPS computes n! and passes the result to continuation k.
func factCPS(n int, k func(int) int) int {
	if n == 0 {
		return k(1)
	}
	return factCPS(n-1, func(r int) int { return k(n * r) })
}

func main() {
	identity := func(x int) int { return x }
	fmt.Println(factCPS(5, identity))
}
