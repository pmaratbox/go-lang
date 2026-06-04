package main

import "fmt"

func pipe(fns ...func(int) int) func(int) int {
	return func(x int) int {
		for _, f := range fns {
			x = f(x)
		}
		return x
	}
}

func main() {
	inc := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	neg := func(x int) int { return -x }
	p := pipe(inc, double, neg)
	fmt.Println(p(3))
}
