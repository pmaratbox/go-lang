package main

import "fmt"

func counted(fn func()) (func(), *int) {
	count := 0
	return func() {
		count++
		fn()
	}, &count
}

func main() {
	wrapped, calls := counted(func() {})
	for i := 0; i < 5; i++ {
		wrapped()
	}
	fmt.Printf("calls: %d\n", *calls)
}
