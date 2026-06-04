package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func partial(f func(int, int) int, a int) func(int) int {
	return func(b int) int {
		return f(a, b)
	}
}

func main() {
	add10 := partial(add, 10)
	fmt.Println(add10(3))
}
