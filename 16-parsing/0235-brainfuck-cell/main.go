package main

import "fmt"

func main() {
	prog := "+++"
	cell := 0
	for _, ch := range prog {
		switch ch {
		case '+':
			cell++
		case '-':
			cell--
		}
	}
	fmt.Println(cell)
}
