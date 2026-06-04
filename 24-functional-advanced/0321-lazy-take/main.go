package main

import (
	"fmt"
	"strings"
)

// naturals lazily emits 1, 2, 3, ... on a channel.
func naturals() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func take(ch <-chan int, n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, <-ch)
	}
	return out
}

func main() {
	nums := take(naturals(), 5)
	parts := make([]string, len(nums))
	for i, v := range nums {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
