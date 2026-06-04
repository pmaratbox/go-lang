package main

import "fmt"

func main() {
	tests := []func() bool{
		func() bool { return 1+1 == 2 },
		func() bool { return len("go") == 2 },
		func() bool { return 10/2 == 5 },
	}

	passed, failed := 0, 0
	for _, t := range tests {
		if t() {
			passed++
		} else {
			failed++
		}
	}

	fmt.Printf("%d passed, %d failed\n", passed, failed)
}
