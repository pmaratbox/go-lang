package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("3 + 4 = %d\n", add(3, 4))
}
