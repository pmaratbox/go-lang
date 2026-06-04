package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countdown(n int, acc []string) []string {
	if n == 0 {
		return acc
	}
	acc = append(acc, strconv.Itoa(n))
	return countdown(n-1, acc)
}

func main() {
	fmt.Println(strings.Join(countdown(5, nil), " "))
}
