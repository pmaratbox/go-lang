package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b := 0, 1
	nums := []string{}
	for i := 0; i < 7; i++ {
		nums = append(nums, fmt.Sprint(a))
		a, b = b, a+b
	}
	fmt.Println(strings.Join(nums, " "))
}
