package main

import (
	"fmt"
	"time"
)

func main() {
	// Fixed UTC instant — never the current time.
	d, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2026-06-15T00:00:00Z")
	fmt.Println(d.Unix())
}
