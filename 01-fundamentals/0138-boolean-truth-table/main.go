package main

import "fmt"

func main() {
	rows := [][2]bool{{true, true}, {true, false}, {false, true}, {false, false}}
	for _, r := range rows {
		a, b := r[0], r[1]
		fmt.Printf("%t %t %t %t %t\n", a, b, a && b, a || b, a != b)
	}
}
