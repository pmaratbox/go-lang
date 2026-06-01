package main

import "fmt"

func task(x int, ch chan<- int) {
	ch <- x
}

func main() {
	ch := make(chan int, 2)
	go task(1, ch)
	go task(2, ch)
	sum := <-ch + <-ch
	fmt.Printf("sum: %d\n", sum)
}
