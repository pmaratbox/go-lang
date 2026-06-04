package main

import "fmt"

type pred func(int) bool

func and(p, q pred) pred {
	return func(n int) bool {
		return p(n) && q(n)
	}
}

func isEven(n int) bool     { return n%2 == 0 }
func isPositive(n int) bool { return n > 0 }

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	test := and(isEven, isPositive)
	fmt.Println(label(test(4)), label(test(-4)))
}
