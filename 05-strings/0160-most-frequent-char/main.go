package main

import "fmt"

func main() {
	s := "hello"
	counts := map[rune]int{}
	var best rune
	bestCount := 0
	for _, r := range s {
		counts[r]++
		if counts[r] > bestCount {
			bestCount = counts[r]
			best = r
		}
	}
	fmt.Println(string(best))
}
