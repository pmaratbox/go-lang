package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 60
	var factors []string
	for n%2 == 0 {
		factors = append(factors, "2")
		n /= 2
	}
	for d := 3; d*d <= n; d += 2 {
		for n%d == 0 {
			factors = append(factors, fmt.Sprint(d))
			n /= d
		}
	}
	if n > 1 {
		factors = append(factors, fmt.Sprint(n))
	}
	fmt.Println(strings.Join(factors, " "))
}
