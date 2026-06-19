package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func g(c *lru.Cache[string, int], k string) string {
	if v, ok := c.Get(k); ok {
		return fmt.Sprint(v)
	}
	return "miss"
}

func main() {
	c, _ := lru.New[string, int](3)
	fmt.Println(g(c, "x"))
}
