package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	seen := map[int]int{}
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			fmt.Printf("%d %d\n", j, i)
			break
		}
		seen[n] = i
	}
}
