package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	targets := []int{2, 0}
	floor := 0
	floors := []int{floor}
	for _, t := range targets {
		for floor != t {
			if floor < t {
				floor++
			} else {
				floor--
			}
			floors = append(floors, floor)
		}
	}
	parts := make([]string, len(floors))
	for i, f := range floors {
		parts[i] = strconv.Itoa(f)
	}
	fmt.Println(strings.Join(parts, " "))
}
