package main

import (
	"fmt"
	"sort"
)

func sortedKey(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat"}
	groups := make(map[string][]string)
	for _, w := range words {
		k := sortedKey(w)
		groups[k] = append(groups[k], w)
	}
	fmt.Println(len(groups))
}
