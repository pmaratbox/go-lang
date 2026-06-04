package main

import "fmt"

type handler struct {
	level int
	next  *handler
}

func (h *handler) handle(level int) {
	if h.level == level {
		fmt.Printf("handled by %d\n", h.level)
		return
	}
	if h.next != nil {
		h.next.handle(level)
	}
}

func main() {
	h3 := &handler{level: 3}
	h2 := &handler{level: 2, next: h3}
	h1 := &handler{level: 1, next: h2}
	h1.handle(2)
}
