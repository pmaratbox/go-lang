package main

import (
	"fmt"
	"strings"
)

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	years := []int{2000, 1900, 2024}
	out := make([]string, len(years))
	for i, y := range years {
		if isLeap(y) {
			out[i] = "yes"
		} else {
			out[i] = "no"
		}
	}
	fmt.Println(strings.Join(out, " "))
}
