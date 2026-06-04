package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 5, 5, 7}
	lower := sort.SearchInts(a, 5)
	upper := sort.Search(len(a), func(i int) bool { return a[i] > 5 })
	fmt.Println(lower, upper)
}
