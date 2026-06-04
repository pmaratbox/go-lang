package main

import "fmt"

func main() {
	months := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	month, day := 3, 1
	doy := day
	for i := 0; i < month-1; i++ {
		doy += months[i]
	}
	fmt.Println(doy)
}
