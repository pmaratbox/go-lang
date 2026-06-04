package main

import (
	"fmt"
	"strings"
)

type Color int

const (
	RED Color = iota
	GREEN
	BLUE
)

func (c Color) String() string {
	return [...]string{"RED", "GREEN", "BLUE"}[c]
}

func main() {
	var names []string
	for c := RED; c <= BLUE; c++ {
		names = append(names, c.String())
	}
	fmt.Println(strings.Join(names, " "))
}
