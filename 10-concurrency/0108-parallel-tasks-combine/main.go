package main

import "fmt"

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)
	go func() { a <- 10 }()
	go func() { b <- 20 }()
	fmt.Println(<-a + <-b)
}
