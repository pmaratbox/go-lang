package main

import "fmt"

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	fmt.Printf("green: %d\n", Green)
	fmt.Printf("blue: %d\n", Blue)
}
