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
		return errors.New("always fails")
	}
	// Allow up to 4 zero-delay retries => 5 total attempts.
	// The op always fails, so the library exhausts every attempt
	// and the closure counter records the total number of calls.
	backoff.Retry(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 4))
	fmt.Println(attempts)
}
