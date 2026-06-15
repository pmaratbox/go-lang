package main

import (
	"fmt"
	"time"
)

func main() {
	// Parse two FIXED ISO dates with the stdlib time library (never time.Now()).
	start, err := time.Parse("2006-01-02", "2026-06-15")
	if err != nil {
		panic(err)
	}
	end, err := time.Parse("2006-01-02", "2026-07-15")
	if err != nil {
		panic(err)
	}
	// Let the library compute the difference and convert to whole days.
	days := int(end.Sub(start).Hours() / 24)
	fmt.Println(days)
}
