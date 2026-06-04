package main

import "fmt"

// sum recursively forks the range [lo, hi) into halves computed concurrently.
func sum(nums []int) int {
	if len(nums) <= 1 {
		if len(nums) == 0 {
			return 0
		}
		return nums[0]
	}
	mid := len(nums) / 2
	left := make(chan int, 1)
	go func() { left <- sum(nums[:mid]) }()
	right := sum(nums[mid:])
	return <-left + right
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(nums))
}
