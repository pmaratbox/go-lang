package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	sum := lo.Reduce([]int{1, 2, 3, 4, 5}, func(acc, x, _ int) int { return acc + x }, 0)
	fmt.Println(sum)
}
