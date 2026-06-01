package main

import (
	"cmp"
	"fmt"
)

func largest[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largest(3, 9))
	fmt.Println(largest("apple", "pear"))
}
