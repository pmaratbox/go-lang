package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	var parts []string
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			parts = append(parts, strconv.Itoa(a[i]))
		}
		if i < len(b) {
			parts = append(parts, strconv.Itoa(b[i]))
		}
	}
	fmt.Println(strings.Join(parts, " "))
}
