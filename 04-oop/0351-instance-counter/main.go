package main

import "fmt"

var instanceCount int

type Widget struct{}

// newWidget acts as the constructor, bumping the class-level counter.
func newWidget() Widget {
	instanceCount++
	return Widget{}
}

func main() {
	newWidget()
	newWidget()
	newWidget()
	fmt.Println(instanceCount)
}
