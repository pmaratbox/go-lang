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
		if attempts < 4 {
			return errors.New("transient failure")
		}
		return nil // succeeds on the fourth call
	}
	// Exponential backoff strategy with a zero base delay so the growing
	// intervals collapse to no waiting. The scripted operation fails three
	// times then succeeds, so the library retries three times and the
	// closure counter records four total attempts.
	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = 0
	exp.MaxInterval = 0
	backoff.Retry(op, exp)
	fmt.Println(attempts)
}
