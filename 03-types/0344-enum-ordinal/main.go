package main

import "fmt"

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

var names = []string{"N", "E", "S", "W"}

func (d Direction) String() string { return names[d] }

func main() {
	fmt.Println(int(S), Direction(3))
}
