package main

import (
	"fmt"
	"time"
)

func main() {
	// Fixed instant — never the current time.
	start, _ := time.Parse("2006-01-02T15:04", "2026-06-15T10:00")
	// Add 90 minutes via the time library.
	end := start.Add(90 * time.Minute)
	// Format as HH:mm.
	fmt.Println(end.Format("15:04"))
}
