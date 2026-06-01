package main

import "fmt"

func inc(x int) int   { return x + 1 }
func twice(x int) int { return x * 2 }

func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	fmt.Println(compose(inc, twice)(3))
}
