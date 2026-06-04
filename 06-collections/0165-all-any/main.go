package main

import "fmt"

func allOf(xs []int, pred func(int) bool) bool {
	for _, x := range xs {
		if !pred(x) {
			return false
		}
	}
	return true
}

func anyOf(xs []int, pred func(int) bool) bool {
	for _, x := range xs {
		if pred(x) {
			return true
		}
	}
	return false
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	xs := []int{2, 4, 6}
	even := func(x int) bool { return x%2 == 0 }
	odd := func(x int) bool { return x%2 != 0 }
	fmt.Println(yesNo(allOf(xs, even)), yesNo(anyOf(xs, odd)))
}
