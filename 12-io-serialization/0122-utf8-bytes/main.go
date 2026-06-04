package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hi"
	parts := make([]string, 0, len(s))
	for _, b := range []byte(s) {
		parts = append(parts, strconv.Itoa(int(b)))
	}
	fmt.Println(strings.Join(parts, " "))
}
