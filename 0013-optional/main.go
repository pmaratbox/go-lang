package main

import "fmt"

func orElse(p *int, fallback int) int {
	if p != nil {
		return *p
	}
	return fallback
}

func main() {
	value := 42
	present := &value
	var absent *int

	fmt.Printf("present: %d\n", orElse(present, -1))
	fmt.Printf("absent: %d\n", orElse(absent, -1))
}
