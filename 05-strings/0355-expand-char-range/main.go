package main

import (
	"fmt"
	"strings"
)

func expand(spec string) string {
	start, end := spec[0], spec[2]
	var b strings.Builder
	for c := start; c <= end; c++ {
		b.WriteByte(c)
	}
	return b.String()
}

func main() {
	fmt.Println(expand("a-e"))
}
