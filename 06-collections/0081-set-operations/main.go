package main

import (
	"fmt"
	"sort"
	"strings"
)

func join(xs []int) string {
	parts := []string{}
	for _, x := range xs {
		parts = append(parts, fmt.Sprint(x))
	}
	return strings.Join(parts, " ")
}

func main() {
	a := map[int]bool{1: true, 2: true, 3: true}
	b := map[int]bool{2: true, 3: true, 4: true}

	unionSet := map[int]bool{}
	for k := range a {
		unionSet[k] = true
	}
	for k := range b {
		unionSet[k] = true
	}
	union := []int{}
	for k := range unionSet {
		union = append(union, k)
	}
	sort.Ints(union)

	inter := []int{}
	for k := range a {
		if b[k] {
			inter = append(inter, k)
		}
	}
	sort.Ints(inter)

	fmt.Println(join(union))
	fmt.Println(join(inter))
}
