package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a, b := 4, 6
	fmt.Println(a / gcd(a, b) * b)
}
