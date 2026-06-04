package main

import "fmt"

func main() {
	n := 1234
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	fmt.Println(rev)
}
