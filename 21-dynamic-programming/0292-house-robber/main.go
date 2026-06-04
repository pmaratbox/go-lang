package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	prev, cur := 0, 0
	for _, n := range nums {
		take := prev + n
		if cur > take {
			take = cur
		}
		prev, cur = cur, take
	}
	fmt.Println(cur)
}
