package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := []string{}
	i := 1
	for {
		parts = append(parts, fmt.Sprint(i))
		i++
		if i > 3 {
			break
		}
	}
	fmt.Println(strings.Join(parts, " "))
}
