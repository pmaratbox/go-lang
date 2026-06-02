package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := strings.Split("a,b,c", ",")
	for i, p := range parts {
		parts[i] = strings.ToUpper(p)
	}
	fmt.Println(strings.Join(parts, "-"))
}
