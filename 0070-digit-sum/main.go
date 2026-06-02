package main

import "fmt"

func main() {
	n := 1234
	total := 0
	for n > 0 {
		total += n % 10
		n /= 10
	}
	fmt.Println(total)
}
