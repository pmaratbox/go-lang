package main

import "fmt"

func inc(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func apply(f func(int) int, x int) int {
	return f(x)
}

func main() {
	fmt.Printf("inc 5 = %d\n", apply(inc, 5))
	fmt.Printf("double 5 = %d\n", apply(double, 5))
}
