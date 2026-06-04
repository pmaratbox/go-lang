package main

import (
	"fmt"
	"strings"
)

func main() {
	rows := 4
	row := []int{1}
	var lines []string
	for r := 0; r < rows; r++ {
		parts := make([]string, len(row))
		for i, v := range row {
			parts[i] = fmt.Sprint(v)
		}
		lines = append(lines, strings.Join(parts, " "))
		next := make([]int, len(row)+1)
		next[0] = 1
		next[len(row)] = 1
		for i := 1; i < len(row); i++ {
			next[i] = row[i-1] + row[i]
		}
		row = next
	}
	fmt.Println(strings.Join(lines, "\n"))
}
