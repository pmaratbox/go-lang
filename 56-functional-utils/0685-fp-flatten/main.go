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
	flat := lo.Flatten([][]int{{1, 2}, {3, 4}, {5, 6}})
	fmt.Println(ij(flat, ","))
}
