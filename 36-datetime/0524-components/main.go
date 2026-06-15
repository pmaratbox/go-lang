package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.Parse("2006-01-02", "2026-06-15")
	fmt.Println(d.Year())
	fmt.Println(int(d.Month()))
	fmt.Println(d.Day())
}
