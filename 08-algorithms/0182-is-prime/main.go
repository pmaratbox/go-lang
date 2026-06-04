package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(label(isPrime(7)), label(isPrime(9)))
}
