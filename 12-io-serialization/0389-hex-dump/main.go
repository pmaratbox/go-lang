package main

import "fmt"

func main() {
	data := []byte("Hi")

	for i, b := range data {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x", b)
	}
	fmt.Println()
}
