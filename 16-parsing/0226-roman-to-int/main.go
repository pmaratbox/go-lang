package main

import "fmt"

func main() {
	roman := "XIV"
	vals := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	for i := 0; i < len(roman); i++ {
		v := vals[roman[i]]
		if i+1 < len(roman) && v < vals[roman[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	fmt.Println(total)
}
