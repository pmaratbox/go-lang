package main

import "fmt"

func main() {
	for _, c := range []int{100, 0} {
		f := c*9/5 + 32
		fmt.Printf("%dC = %dF\n", c, f)
	}
}
