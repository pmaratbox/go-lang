package main

import "fmt"

func main() {
	a, b := 6, 3
	fmt.Println("and:", a&b)
	fmt.Println("or:", a|b)
	fmt.Println("xor:", a^b)
	fmt.Println("shift:", a<<1)
}
