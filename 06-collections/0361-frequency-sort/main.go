package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1, 1, 2, 3, 3, 3}

	counts := make(map[int]int)
	var order []int
	for _, n := range nums {
		if _, seen := counts[n]; !seen {
			order = append(order, n)
		}
		counts[n]++
	}

	sort.SliceStable(order, func(i, j int) bool {
		return counts[order[i]] > counts[order[j]]
	})

	var parts []string
	for _, n := range order {
		for k := 0; k < counts[n]; k++ {
			parts = append(parts, strconv.Itoa(n))
		}
	}
	fmt.Println(strings.Join(parts, " "))
}
