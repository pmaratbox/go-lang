package main

import "fmt"

// Bit flags declared with iota-style shifts.
const (
	READ  = 1 << iota // 1
	WRITE             // 2
)

func main() {
	flags := READ | WRITE

	set := "no"
	if flags&WRITE != 0 {
		set = "yes"
	}

	fmt.Printf("%d %s\n", flags, set)
}
