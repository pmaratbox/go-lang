package main

import "fmt"

func main() {
	fmt.Println("open")
	defer fmt.Println("close")
	fmt.Println("use")
}
