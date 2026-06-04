package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
