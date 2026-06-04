package main

import (
	"fmt"
	"strings"
)

func main() {
	adj := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2},
	}
	visited := map[int]bool{0: true}
	queue := []int{0}
	var order []int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)
		for _, n := range adj[cur] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	var parts []string
	for _, v := range order {
		parts = append(parts, fmt.Sprint(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
