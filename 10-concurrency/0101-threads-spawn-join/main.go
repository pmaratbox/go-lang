package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const n = 3
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = i * i // trivial work
		}()
	}
	wg.Wait()
	fmt.Printf("done: %d\n", n)
}
