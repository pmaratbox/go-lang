package main

import "fmt"

// op succeeds only on attempt 3.
func op(attempt int) bool {
	return attempt == 3
}

func main() {
	const maxAttempts = 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if op(attempt) {
			fmt.Printf("ok after %d\n", attempt)
			break
		}
	}
}
