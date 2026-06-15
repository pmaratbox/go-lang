package main

import (
	"fmt"

	"github.com/vektah/goparsify"
)

func main() {
	// Sequence combinator: parse char 'a' THEN char 'b', then combine
	// the two captured tokens into a single string via Map.
	parser := goparsify.Seq("a", "b").Map(func(n *goparsify.Result) {
		n.Result = n.Child[0].Token + n.Child[1].Token
	})

	result, err := goparsify.Run(parser, "ab")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
