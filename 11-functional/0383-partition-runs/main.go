package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1, 1, 2, 3, 3, 3}
	var runs []string
	var cur []string
	for i, n := range nums {
		if i > 0 && n != nums[i-1] {
			runs = append(runs, strings.Join(cur, " "))
			cur = nil
		}
		cur = append(cur, strconv.Itoa(n))
	}
	if len(cur) > 0 {
		runs = append(runs, strings.Join(cur, " "))
	}
	fmt.Println(strings.Join(runs, "|"))
}
