package main

import "fmt"

func main() {
	n := 90210
	count := 0
	for x := n; x != 0; x /= 10 {
		count++
	}
	fmt.Println(count)
}
