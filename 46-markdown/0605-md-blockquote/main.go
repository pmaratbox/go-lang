package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
)

func render(src string) string {
	var b bytes.Buffer
	goldmark.Convert([]byte(src), &b)
	return strings.TrimRight(b.String(), "\n")
}

func main() {
	fmt.Println(render("> quote"))
}
