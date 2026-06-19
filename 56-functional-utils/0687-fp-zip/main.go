package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func main() {
	z := lo.Zip2([]int{1, 2, 3}, []string{"a", "b", "c"})
	pairs := lo.Map(z, func(t lo.Tuple2[int, string], _ int) string {
		return fmt.Sprintf("%d%s", t.A, t.B)
	})
	fmt.Println(strings.Join(pairs, ","))
}
