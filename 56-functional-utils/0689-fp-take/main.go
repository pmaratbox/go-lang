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
	fmt.Println(ij(lo.Subset([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 3), ","))
}
