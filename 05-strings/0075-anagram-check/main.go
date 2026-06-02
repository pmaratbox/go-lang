package main

import (
	"fmt"
	"sort"
)

func sortedString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func main() {
	pairs := [][2]string{{"listen", "silent"}, {"hello", "world"}}
	for _, p := range pairs {
		ans := "no"
		if sortedString(p[0]) == sortedString(p[1]) {
			ans = "yes"
		}
		fmt.Printf("%s/%s: %s\n", p[0], p[1], ans)
	}
}
