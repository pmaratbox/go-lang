package main

import (
	"fmt"
	"time"
)

func main() {
	// Parse a FIXED UTC instant with the stdlib time library (never time.Now()).
	instant, err := time.Parse("2006-01-02T15:04:05Z07:00", "2026-06-15T12:00:00Z")
	if err != nil {
		panic(err)
	}
	// Build a FIXED +05:00 offset zone (no named tz / OS tzdata) and convert.
	plus5 := time.FixedZone("", 5*3600)
	local := instant.In(plus5)
	// The library computes the local hour: 12 UTC + 5 = 17.
	fmt.Println(local.Hour())
}
