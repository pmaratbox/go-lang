package main

import (
	"fmt"
	"strings"
)

func shellSort(a []int) {
	n := len(a)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			tmp := a[i]
			j := i
			for ; j >= gap && a[j-gap] > tmp; j -= gap {
				a[j] = a[j-gap]
			}
			a[j] = tmp
		}
	}
}

func main() {
	a := []int{5, 2, 8, 1, 9, 3}
	shellSort(a)
	parts := make([]string, len(a))
	for i, v := range a {
		parts[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
