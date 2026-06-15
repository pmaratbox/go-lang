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

	// A '+'-separated sequence of integers, built from Seq + Many combinators.
	rest := goparsify.Many(goparsify.Seq("+", num).Map(func(n *goparsify.Result) {
		n.Result = n.Child[1].Result
	}))

	expr := goparsify.Seq(num, rest).Map(func(n *goparsify.Result) {
		sum := n.Child[0].Result.(int)
		for _, c := range n.Child[1].Child {
			sum += c.Result.(int)
		}
		n.Result = sum
	})

	result, err := goparsify.Run(expr, "10+20+30")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
