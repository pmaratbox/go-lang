package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	num, den := 6, 8
	g := gcd(num, den)
	fmt.Printf("%d/%d\n", num/g, den/g)
}
