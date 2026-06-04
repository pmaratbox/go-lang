package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "[1,2,3]"
	inner := strings.Trim(text, "[]")
	sum := 0
	for _, part := range strings.Split(inner, ",") {
		n, _ := strconv.Atoi(strings.TrimSpace(part))
		sum += n
	}
	fmt.Println(sum)
}
