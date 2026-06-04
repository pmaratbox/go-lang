package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	left := map[string]int{"a": 1, "b": 2}
	right := map[string]int{"b": 3, "c": 4}

	merged := make(map[string]int)
	for k, v := range left {
		merged[k] = v
	}
	for k, v := range right {
		merged[k] = v
	}

	keys := make([]string, 0, len(merged))
	for k := range merged {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = fmt.Sprintf("%s:%d", k, merged[k])
	}
	fmt.Println(strings.Join(parts, " "))
}
