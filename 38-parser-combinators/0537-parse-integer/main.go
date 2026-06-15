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
	result, err := goparsify.Run(num, "42")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
