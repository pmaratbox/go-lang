package main

import "fmt"

func main() {
	m := [2][2]int{{1, 2}, {3, 4}}
	v := [2]int{5, 6}
	var out [2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			out[i] += m[i][j] * v[j]
		}
	}
	fmt.Printf("%d %d\n", out[0], out[1])
}
