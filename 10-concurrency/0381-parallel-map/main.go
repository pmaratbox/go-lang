package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	input := []int{1, 2, 3, 4}
	results := make([]int, len(input))
	var wg sync.WaitGroup
	for i, v := range input {
		wg.Add(1)
		go func(idx, val int) {
			defer wg.Done()
			results[idx] = val * val
		}(i, v)
	}
	wg.Wait()

	parts := make([]string, len(results))
	for i, r := range results {
		parts[i] = fmt.Sprint(r)
	}
	fmt.Println(strings.Join(parts, " "))
}
