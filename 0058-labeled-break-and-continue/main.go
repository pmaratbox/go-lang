package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > i {
				continue outer
			}
			if i*j == 4 {
				fmt.Printf("stop at %d,%d\n", i, j)
				break outer
			}
			fmt.Printf("%d,%d\n", i, j)
		}
	}
}
