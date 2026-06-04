package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	const workers = 10
	const perWorker = 10
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < perWorker; j++ {
				for {
					old := atomic.LoadInt64(&counter)
					if atomic.CompareAndSwapInt64(&counter, old, old+1) {
						break
					}
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
