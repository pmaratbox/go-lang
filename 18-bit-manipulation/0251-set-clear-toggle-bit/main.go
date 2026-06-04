package main

import "fmt"

func main() {
	const pos = 1
	set := 0 | (1 << pos)
	clear := 2 &^ (1 << pos)
	toggle := 0 ^ (1 << pos)
	fmt.Println(set, clear, toggle)
}
