package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"one", "two", "three"}
	groups := map[int][]string{}
	for _, w := range words {
		groups[len(w)] = append(groups[len(w)], w)
	}

	keys := []int{}
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	parts := []string{}
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%d:[%s]", k, strings.Join(groups[k], ",")))
	}
	fmt.Println(strings.Join(parts, " "))
}
