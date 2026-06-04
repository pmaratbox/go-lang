package main

import "fmt"

func main() {
	defer fmt.Print("1")
	defer fmt.Print("2 ")
	defer fmt.Print("3 ")
}
