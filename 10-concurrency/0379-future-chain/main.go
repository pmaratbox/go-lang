package main

import "fmt"

// future runs fn in a goroutine and returns a channel delivering its result.
func future(fn func() int) <-chan int {
	ch := make(chan int, 1)
	go func() { ch <- fn() }()
	return ch
}

func main() {
	a := future(func() int { return 5 })
	b := future(func() int { return <-a * 2 })
	c := future(func() int { return <-b + 1 })
	fmt.Println(<-c)
}
