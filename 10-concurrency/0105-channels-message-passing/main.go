package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ch := make(chan int)
	go func() {
		for _, v := range []int{1, 2, 3} {
			ch <- v
		}
		close(ch)
	}()

	var parts []string
	for v := range ch {
		parts = append(parts, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(parts, " "))
}
