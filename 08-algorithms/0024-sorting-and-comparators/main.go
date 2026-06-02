package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func join(xs []int) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = strconv.Itoa(x)
	}
	return strings.Join(parts, " ")
}

func main() {
	asc := []int{3, 1, 2}
	slices.Sort(asc)
	desc := []int{3, 1, 2}
	slices.SortFunc(desc, func(a, b int) int { return cmp.Compare(b, a) })
	fmt.Printf("asc: %s\n", join(asc))
	fmt.Printf("desc: %s\n", join(desc))
}
