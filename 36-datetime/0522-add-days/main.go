package main

import (
	"fmt"
	"time"
)

func main() {
	// Fixed date — never the current time.
	d, _ := time.Parse("2006-01-02", "2026-06-15")

	// Add 10 days via the library's date arithmetic.
	result := d.AddDate(0, 0, 10)

	fmt.Println(result.Format("2006-01-02"))
}
