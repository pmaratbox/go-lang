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
	doubled := lo.Map([]int{1, 2, 3}, func(x int, _ int) int { return x * 2 })
	fmt.Println(ij(doubled, ","))
}
