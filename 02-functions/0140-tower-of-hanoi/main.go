package main

import "fmt"

func moves(n int) int {
	if n == 0 {
		return 0
	}
	return 2*moves(n-1) + 1
}

func main() {
	fmt.Println(moves(3))
}
