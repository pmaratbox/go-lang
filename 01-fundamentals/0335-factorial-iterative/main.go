package main

import "fmt"

func main() {
	result := 1
	for i := 2; i <= 5; i++ {
		result *= i
	}
	fmt.Println(result)
}
