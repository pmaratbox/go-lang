package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	n1, d1 := 1, 2
	n2, d2 := 1, 3
	num := n1*d2 + n2*d1
	den := d1 * d2
	g := gcd(num, den)
	fmt.Printf("%d/%d\n", num/g, den/g)
}
