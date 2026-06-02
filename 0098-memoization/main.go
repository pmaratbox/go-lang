package main

import "fmt"

var cache = map[int]int{}

func fib(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = fib(n-1) + fib(n-2)
	return cache[n]
}

func main() {
	fmt.Println(fib(10))
}
