package main

import "fmt"

func main() {
	const price = 25
	total := 0
	dispensed := false
	for _, coin := range []int{10, 10, 5} {
		total += coin
		if total >= price {
			dispensed = true
			break
		}
	}
	if dispensed {
		fmt.Println("dispensed")
	}
}
