package main

import "fmt"

func isPerfect(n int) bool {
	sum := 0
	for d := 1; d < n; d++ {
		if n%d == 0 {
			sum += d
		}
	}
	return sum == n
}

func label(n int) string {
	if isPerfect(n) {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Printf("%s %s\n", label(6), label(8))
}
