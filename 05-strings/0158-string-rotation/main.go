package main

import (
	"fmt"
	"strings"
)

func isRotation(a, b string) bool {
	return len(a) == len(b) && strings.Contains(a+a, b)
}

func main() {
	if isRotation("abcd", "cdab") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
