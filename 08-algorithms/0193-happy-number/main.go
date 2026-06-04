package main

import "fmt"

func sumSquares(n int) int {
	s := 0
	for n > 0 {
		d := n % 10
		s += d * d
		n /= 10
	}
	return s
}

func isHappy(n int) bool {
	seen := map[int]bool{}
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumSquares(n)
	}
	return n == 1
}

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(label(isHappy(19)))
}
