package main

import "fmt"

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	cur, best := a[0], a[0]
	for _, x := range a[1:] {
		if cur+x > x {
			cur = cur + x
		} else {
			cur = x
		}
		if cur > best {
			best = cur
		}
	}
	fmt.Println(best)
}
