package main

import "fmt"

func binomial(n, k int) int {
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}

func main() {
	fmt.Println(binomial(5, 2))
}
