package main

import "fmt"

func inc(p *int) {
	*p++
}

func main() {
	n := 1
	fmt.Printf("before: %d\n", n)
	inc(&n)
	fmt.Printf("after: %d\n", n)
}
