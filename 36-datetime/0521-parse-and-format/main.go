package main

import (
	"fmt"
	"time"
)

func main() {
	// Parse a FIXED ISO date with the stdlib time library (never time.Now()).
	d, err := time.Parse("2006-01-02", "2026-06-15")
	if err != nil {
		panic(err)
	}
	// Format it back to ISO yyyy-MM-dd using the library.
	fmt.Println(d.Format("2006-01-02"))
}
