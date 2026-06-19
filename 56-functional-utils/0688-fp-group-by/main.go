package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ij(xs []int, sep string) string {
	return strings.Join(lo.Map(xs, func(x int, _ int) string { return strconv.Itoa(x) }), sep)
}

func main() {
	g := lo.GroupBy([]int{1, 2, 3, 4, 5, 6}, func(x int) string {
		if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	keys := lo.Keys(g)
	sort.Strings(keys)
	fmt.Println(strings.Join(lo.Map(keys, func(k string, _ int) string {
		return k + ":" + ij(g[k], ",")
	}), ";"))
}
