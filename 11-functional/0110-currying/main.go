package main

import "fmt"

func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	fmt.Println(add(2)(3))
}
