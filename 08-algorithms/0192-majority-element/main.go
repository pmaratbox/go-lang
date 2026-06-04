package main

import "fmt"

func main() {
	a := []int{2, 2, 1, 2, 3, 2}
	candidate, count := 0, 0
	for _, x := range a {
		if count == 0 {
			candidate = x
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	fmt.Println(candidate)
}
