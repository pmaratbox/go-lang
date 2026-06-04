package main

import "fmt"

func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func inc(n int) int {
	return n + 1
}

func main() {
	fmt.Println(applyTwice(inc, 3))
}
