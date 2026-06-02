package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 10
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	parts := []string{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			parts = append(parts, fmt.Sprint(i))
		}
	}
	fmt.Println(strings.Join(parts, " "))
}
