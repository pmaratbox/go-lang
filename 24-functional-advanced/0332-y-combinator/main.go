package main

import "fmt"

type fn func(int) int

// fix is a fixed-point combinator: it ties the recursive knot for gen
// without gen referring to itself by name.
func fix(gen func(fn) fn) fn {
	var self fn
	self = func(n int) int { return gen(func(m int) int { return self(m) })(n) }
	return self
}

func main() {
	factGen := func(rec fn) fn {
		return func(n int) int {
			if n == 0 {
				return 1
			}
			return n * rec(n-1)
		}
	}
	factorial := fix(factGen)
	fmt.Println(factorial(5))
}
