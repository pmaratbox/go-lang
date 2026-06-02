package main

import (
	"fmt"
	"strings"
)

func main() {
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{5, 6}, {7, 8}}
	var result [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	for i := 0; i < 2; i++ {
		parts := []string{}
		for j := 0; j < 2; j++ {
			parts = append(parts, fmt.Sprint(result[i][j]))
		}
		fmt.Println(strings.Join(parts, " "))
	}
}
