package main

import "fmt"

func zeller(y, m, d int) int {
	if m < 3 {
		m += 12
		y--
	}
	k := y % 100
	j := y / 100
	h := (d + (13*(m+1))/5 + k + k/4 + j/4 + 5*j) % 7
	return h
}

func main() {
	// Zeller's h: 0=Saturday, 1=Sunday, ... 6=Friday
	names := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	h := zeller(2000, 1, 1)
	fmt.Println(names[h])
}
