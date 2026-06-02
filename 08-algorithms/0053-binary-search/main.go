package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 7

	lo, hi, index := 0, len(nums)-1, -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	fmt.Printf("found %d at index %d\n", target, index)
}
