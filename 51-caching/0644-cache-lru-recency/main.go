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
	c.Add("a", 1)
	c.Add("b", 2)
	c.Add("c", 3)
	c.Get("a")    // promotes "a" to most-recently-used; "b" is now the LRU
	c.Add("d", 4) // overflows capacity 3; evicts least-recently-used "b"
	fmt.Println(g(c, "a"), g(c, "b"))
}
