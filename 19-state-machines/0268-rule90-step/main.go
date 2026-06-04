package main

import "fmt"

func main() {
	row := "00100"
	n := len(row)
	bit := func(i int) int {
		if i < 0 || i >= n {
			return 0
		}
		return int(row[i] - '0')
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = byte('0' + bit(i-1)^bit(i+1))
	}
	fmt.Println(string(out))
}
