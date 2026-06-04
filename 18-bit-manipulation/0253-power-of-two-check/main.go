package main

import "fmt"

func isPow2(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func label(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(label(isPow2(16)), label(isPow2(18)))
}
