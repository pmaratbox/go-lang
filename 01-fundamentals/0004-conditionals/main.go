package main

import "fmt"

func main() {
	n := 7

	if n < 10 {
		fmt.Printf("%d is less than 10\n", n)
	} else if n == 10 {
		fmt.Printf("%d is equal to 10\n", n)
	} else {
		fmt.Printf("%d is greater than 10\n", n)
	}
}
