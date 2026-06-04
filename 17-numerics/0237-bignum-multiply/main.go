package main

import (
	"fmt"
	"strings"
)

func mulStrings(a, b string) string {
	n, m := len(a), len(b)
	prod := make([]int, n+m)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			prod[i+j+1] += int(a[i]-'0') * int(b[j]-'0')
		}
	}
	for k := len(prod) - 1; k > 0; k-- {
		prod[k-1] += prod[k] / 10
		prod[k] %= 10
	}
	var sb strings.Builder
	for _, d := range prod {
		sb.WriteByte(byte('0' + d))
	}
	out := strings.TrimLeft(sb.String(), "0")
	if out == "" {
		return "0"
	}
	return out
}

func main() {
	fmt.Println(mulStrings("123", "456"))
}
