package main

import "fmt"

func main() {
	n := 5
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Println(a)
}
