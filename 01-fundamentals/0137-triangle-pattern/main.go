package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
