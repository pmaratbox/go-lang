package main

import "fmt"

func quickselect(a []int, k int) int {
	lo, hi := 0, len(a)-1
	for {
		pivot := a[hi]
		i := lo
		for j := lo; j < hi; j++ {
			if a[j] < pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[hi] = a[hi], a[i]
		switch {
		case k < i:
			hi = i - 1
		case k > i:
			lo = i + 1
		default:
			return a[i]
		}
	}
}

func main() {
	a := []int{7, 10, 4, 3, 20, 15}
	fmt.Println(quickselect(a, 2))
}
