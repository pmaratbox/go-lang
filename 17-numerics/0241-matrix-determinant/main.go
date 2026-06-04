package main

import "fmt"

func main() {
	m := [2][2]int{{1, 2}, {3, 4}}
	det := m[0][0]*m[1][1] - m[0][1]*m[1][0]
	fmt.Println(det)
}
