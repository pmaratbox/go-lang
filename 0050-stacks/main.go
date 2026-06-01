package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stack := []int{}
	for _, n := range []int{1, 2, 3} {
		stack = append(stack, n)
	}

	var popped []string
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		popped = append(popped, strconv.Itoa(top))
	}

	fmt.Println(strings.Join(popped, " "))
}
