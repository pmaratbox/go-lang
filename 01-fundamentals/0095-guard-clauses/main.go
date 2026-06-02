package main

import "fmt"

func classify(n int) string {
	if n < 0 {
		return "negative"
	}
	if n == 0 {
		return "zero"
	}
	return "positive"
}

func main() {
	for _, n := range []int{-1, 0, 5} {
		fmt.Println(classify(n))
	}
}
