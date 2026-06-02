package main

import (
	"fmt"
	"strings"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for j := 0; j < cols; j++ {
		transposed[j] = make([]int, rows)
		for i := 0; i < rows; i++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	for _, row := range transposed {
		parts := []string{}
		for _, x := range row {
			parts = append(parts, fmt.Sprint(x))
		}
		fmt.Println(strings.Join(parts, " "))
	}
}
