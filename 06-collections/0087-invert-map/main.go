package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	inv := map[int]string{}
	for k, v := range m {
		inv[v] = k
	}
	keys := []int{}
	for k := range inv {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	parts := []string{}
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%d:%s", k, inv[k]))
	}
	fmt.Println(strings.Join(parts, " "))
}
