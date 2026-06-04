package main

import (
	"fmt"
	"strings"
)

func dutchFlag(a []int) {
	low, mid, high := 0, 0, len(a)-1
	for mid <= high {
		switch a[mid] {
		case 0:
			a[low], a[mid] = a[mid], a[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			a[mid], a[high] = a[high], a[mid]
			high--
		}
	}
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	dutchFlag(a)
	parts := make([]string, len(a))
	for i, v := range a {
		parts[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
