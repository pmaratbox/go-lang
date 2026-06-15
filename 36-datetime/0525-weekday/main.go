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
	// time.Weekday is Sunday-based (Sunday=0 .. Saturday=6).
	// Convert to ISO weekday numbering (Monday=1 .. Sunday=7).
	iso := (int(d.Weekday())+6)%7 + 1
	fmt.Println(iso)
}
