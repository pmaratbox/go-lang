package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	items := []string{"a", "b", "a", "c", "b", "a"}
	const k = 2

	counts := make(map[string]int)
	var order []string
	for _, it := range items {
		if _, ok := counts[it]; !ok {
			order = append(order, it)
		}
		counts[it]++
	}

	sort.SliceStable(order, func(i, j int) bool {
		return counts[order[i]] > counts[order[j]]
	})

	top := order
	if len(top) > k {
		top = top[:k]
	}
	fmt.Println(strings.Join(top, " "))
}
