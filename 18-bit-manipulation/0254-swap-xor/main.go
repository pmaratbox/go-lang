package main

import "fmt"

func main() {
	a, b := 3, 5
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
