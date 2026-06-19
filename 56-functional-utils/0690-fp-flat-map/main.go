package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ij(xs []int, sep string) string {
	return strings.Join(lo.Map(xs, func(x int, _ int) string { return strconv.Itoa(x) }), sep)
}

func main() {
	result := lo.FlatMap([]int{1, 2, 3}, func(x int, _ int) []int { return []int{x, x * 10} })
	fmt.Println(ij(result, ","))
}
