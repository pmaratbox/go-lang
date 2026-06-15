package main

import (
	"fmt"

	"github.com/vektah/goparsify"
)

func main() {
	// Alternative / choice: try "cat", otherwise "dog".
	animal := goparsify.Any("cat", "dog").Map(func(n *goparsify.Result) {
		n.Result = n.Token
	})
	result, err := goparsify.Run(animal, "dog")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(result)
}
