package main

import "fmt"

func main() {
	xs := []string{"a", "bbb", "cc"}

	best := xs[0]
	for _, s := range xs[1:] {
		if len(s) > len(best) {
			best = s
		}
	}
	fmt.Println(best)
}
