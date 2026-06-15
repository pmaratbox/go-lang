package main

import (
	"fmt"
	"strconv"

	"github.com/vektah/goparsify"
)

func main() {
	// goparsify automatically consumes leading whitespace before each token,
	// so a plain digit parser already skips the spaces in front of "42".
	num := goparsify.Regex(`[0-9]+`).Map(func(n *goparsify.Result) {
		n.Result, _ = strconv.Atoi(n.Token)
	})

	// goparsify.Run trims any trailing whitespace left after the parser runs,
	// so the surrounding spaces in "  42  " are skipped entirely.
	result, err := goparsify.Run(num, "  42  ")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
