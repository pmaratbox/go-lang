package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	scores := map[string]int{"a": 3, "b": 1, "c": 2}
	keys := []string{}
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return scores[keys[i]] < scores[keys[j]] })
	parts := []string{}
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s:%d", k, scores[k]))
	}
	fmt.Println(strings.Join(parts, " "))
}
