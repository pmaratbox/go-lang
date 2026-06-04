package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 3
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			// arrive at the barrier
			wg.Done()
			wg.Wait()
		}()
	}
	wg.Wait()
	fmt.Println("all reached:", n)
}
