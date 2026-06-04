package main

import (
	"fmt"
	"sync"
)

func main() {
	var r1, r2 int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); r1 = 3 * 3 }()
	go func() { defer wg.Done(); r2 = 4 * 4 }()
	wg.Wait()
	fmt.Println(r1 + r2)
}
