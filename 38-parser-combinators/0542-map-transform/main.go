package main

import (
	"fmt"
	"strconv"

	"github.com/vektah/goparsify"
)

func main() {
	// Parse one-or-more digits, then Map the matched Token into an int
	// doubled via strconv.Atoi, storing the transformed value in Result.Result.
	doubled := goparsify.Regex(`[0-9]+`).Map(func(n *goparsify.Result) {
		v, _ := strconv.Atoi(n.Token)
		n.Result = v * 2
	})

	result, err := goparsify.Run(doubled, "21")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
