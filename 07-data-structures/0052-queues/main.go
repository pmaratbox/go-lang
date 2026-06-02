package main

import (
	"fmt"
	"strings"
)

func main() {
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)

	out := []string{}
	for len(queue) > 0 {
		out = append(out, fmt.Sprint(queue[0]))
		queue = queue[1:]
	}
	fmt.Println(strings.Join(out, " "))
}
