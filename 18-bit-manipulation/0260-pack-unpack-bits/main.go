package main

import "fmt"

func main() {
	r, g, b := 1, 2, 3
	packed := (r << 16) | (g << 8) | b
	fmt.Println((packed>>16)&0xff, (packed>>8)&0xff, packed&0xff)
}
