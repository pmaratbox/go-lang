package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4}
	acc := 1
	var out []string
	for _, n := range nums {
		acc *= n
		out = append(out, strconv.Itoa(acc))
	}
	fmt.Println(strings.Join(out, " "))
}
