package main

import "fmt"

func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func main() {
	for _, n := range []int{4, 3} {
		if isEven(n) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
