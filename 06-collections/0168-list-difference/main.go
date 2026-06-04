package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{2, 4}

	remove := make(map[int]struct{}, len(b))
	for _, v := range b {
		remove[v] = struct{}{}
	}

	var diff []string
	for _, v := range a {
		if _, ok := remove[v]; !ok {
			diff = append(diff, fmt.Sprintf("%d", v))
		}
	}
	fmt.Println(strings.Join(diff, " "))
}
