package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const mask = 5
	var parts []string
	for sub := mask; ; sub = (sub - 1) & mask {
		parts = append(parts, strconv.Itoa(sub))
		if sub == 0 {
			break
		}
	}
	fmt.Println(strings.Join(parts, " "))
}
