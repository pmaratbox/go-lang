package main

import (
	"errors"
	"fmt"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	attempts := 0
	op := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("transient failure")
		}
		return nil // succeeds on the third call
	}
	// Fixed (constant) backoff strategy with a zero delay between attempts.
	// The operation fails twice then succeeds, so the library retries twice
	// and the closure counter records three total attempts.
	fixed := backoff.NewConstantBackOff(0)
	backoff.Retry(op, fixed)
	fmt.Println(attempts)
}
