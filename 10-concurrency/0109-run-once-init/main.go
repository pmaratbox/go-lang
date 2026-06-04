package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var once sync.Once
	var initCount int32

	init := func() {
		atomic.AddInt32(&initCount, 1)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(init)
		}()
	}
	wg.Wait()
	fmt.Printf("init count: %d\n", initCount)
}
