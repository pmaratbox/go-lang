package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parts := make([]string, 0, 4)
	for n := 0; n < 4; n++ {
		parts = append(parts, strconv.Itoa(n^(n>>1)))
	}
	fmt.Println(strings.Join(parts, " "))
}
