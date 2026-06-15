package main

import (
	"fmt"
	"strconv"

	"github.com/vektah/goparsify"
)

func main() {
	num := goparsify.Regex(`[0-9]+`).Map(func(n *goparsify.Result) {
		n.Result, _ = strconv.Atoi(n.Token)
	})

	// sepBy: integers separated by ',' built from Seq + Many combinators.
	rest := goparsify.Many(goparsify.Seq(",", num).Map(func(n *goparsify.Result) {
		n.Result = n.Child[1].Result
	}))

	list := goparsify.Seq(num, rest).Map(func(n *goparsify.Result) {
		nums := []int{n.Child[0].Result.(int)}
		for _, c := range n.Child[1].Child {
			nums = append(nums, c.Result.(int))
		}
		n.Result = nums
	})

	result, err := goparsify.Run(list, "1,2,3")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	sum := 0
	for _, v := range result.([]int) {
		sum += v
	}
	fmt.Println(sum)
}
