package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Printf("count: %d\n", len(nums))
	fmt.Printf("first: %d\n", nums[0])
	fmt.Printf("last: %d\n", nums[len(nums)-1])
}
