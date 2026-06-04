package main

import (
	"fmt"
	"strings"
)

func naturals() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
	}()
	return out
}

func take(ch <-chan int, n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, <-ch)
	}
	return out
}

func main() {
	evens := filter(naturals(), func(x int) bool { return x%2 == 0 })
	nums := take(evens, 3)
	parts := make([]string, len(nums))
	for i, v := range nums {
		parts[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println(strings.Join(parts, " "))
}
