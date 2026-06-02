package main

import "fmt"

func main() {
	for _, n := range []int{4, 7} {
		label := "odd"
		if n%2 == 0 {
			label = "even"
		}
		fmt.Printf("%d: %s\n", n, label)
	}
}
