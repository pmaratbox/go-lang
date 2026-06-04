package main

import "fmt"

type strategy func(a, b int) int

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }

func apply(s strategy, a, b int) int { return s(a, b) }

func main() {
	fmt.Println(apply(add, 3, 4), apply(mul, 3, 4))
}
