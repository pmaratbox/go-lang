package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func main() {
	c, _ := lru.New[string, int](3)
	c.Add("a", 1)
	fmt.Println(c.Contains("a"), c.Contains("x"))
}
