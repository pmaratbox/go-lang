package main

import "fmt"

func main() {
	f := 3.9                     // float64 variable
	truncated := int(f)          // 3 (truncates toward zero)

	i := 3                       // int variable
	widened := float64(i)        // 3.0

	fmt.Printf("%d %.1f\n", truncated, widened)
}
