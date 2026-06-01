package main

import "fmt"

func yn(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	nums := map[int]struct{}{}
	for _, n := range []int{1, 2, 2, 3} {
		nums[n] = struct{}{}
	}
	fmt.Printf("size: %d\n", len(nums))
	_, has2 := nums[2]
	_, has5 := nums[5]
	fmt.Printf("has 2: %s\n", yn(has2))
	fmt.Printf("has 5: %s\n", yn(has5))
}
