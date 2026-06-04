package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key, value int
}

type LruCache struct {
	capacity int
	order    *list.List
	items    map[int]*list.Element
}

func newLruCache(capacity int) *LruCache {
	return &LruCache{
		capacity: capacity,
		order:    list.New(),
		items:    make(map[int]*list.Element),
	}
}

func (c *LruCache) get(key int) int {
	if el, ok := c.items[key]; ok {
		c.order.MoveToFront(el)
		return el.Value.(*entry).value
	}
	return -1
}

func (c *LruCache) put(key, value int) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).value = value
		c.order.MoveToFront(el)
		return
	}
	if c.order.Len() >= c.capacity {
		back := c.order.Back()
		if back != nil {
			c.order.Remove(back)
			delete(c.items, back.Value.(*entry).key)
		}
	}
	c.items[key] = c.order.PushFront(&entry{key, value})
}

func main() {
	c := newLruCache(2)
	c.put(1, 1)
	c.put(2, 2)
	c.get(1)
	c.put(3, 3)
	fmt.Println(c.get(1), c.get(2))
}
