package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1.0
	for {
		next := x - (x*x-2)/(2*x)
		if math.Abs(next-x) < 1e-12 {
			x = next
			break
		}
		x = next
	}
	fmt.Printf("%.4f\n", x)
}
