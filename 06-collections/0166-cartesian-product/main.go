package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []int{1, 2}
	letters := []string{"a", "b"}

	var pairs []string
	for _, n := range nums {
		for _, l := range letters {
			pairs = append(pairs, fmt.Sprintf("%d%s", n, l))
		}
	}
	fmt.Println(strings.Join(pairs, " "))
}
