package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func main() {
	c, _ := lru.New[string, int](3)
	c.Add("a", 1)
	c.Add("b", 2)
	c.Add("c", 3)
	c.Add("d", 4)
	fmt.Println(c.Len())
}
