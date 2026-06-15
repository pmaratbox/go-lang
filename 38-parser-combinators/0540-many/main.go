package main

import (
	"fmt"

	"github.com/vektah/goparsify"
)

func main() {
	// Many combinator: apply the char-'a' parser zero-or-more times.
	// Each successful match becomes a child of the result; Map turns the
	// number of children into the count held in Result.
	parser := goparsify.Many("a").Map(func(n *goparsify.Result) {
		n.Result = len(n.Child)
	})

	result, err := goparsify.Run(parser, "aaaa")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
