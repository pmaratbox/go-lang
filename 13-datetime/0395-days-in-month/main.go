package main

import "fmt"

func febDays(year int) int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return 29
	}
	return 28
}

func main() {
	fmt.Printf("%d %d\n", febDays(2000), febDays(2001))
}
