package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder
	for i, n := range []int{1, 2, 3} {
		if i > 0 {
			sb.WriteString("-")
		}
		fmt.Fprintf(&sb, "%d", n)
	}
	fmt.Println(sb.String())
}
