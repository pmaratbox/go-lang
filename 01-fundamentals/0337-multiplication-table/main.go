package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 3; i++ {
		cells := make([]string, 0, 3)
		for j := 1; j <= 3; j++ {
			cells = append(cells, strconv.Itoa(i*j))
		}
		fmt.Println(strings.Join(cells, " "))
	}
}
