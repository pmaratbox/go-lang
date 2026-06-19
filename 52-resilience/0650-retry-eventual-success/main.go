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
		if attempts < 2 {
			return errors.New("transient failure")
		}
		return nil // succeeds on the second call
	}
	// Allow up to 4 retries with zero delay. The op fails once and then
	// succeeds, so the library retries exactly once: 2 attempts total.
	backoff.Retry(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 4))
	fmt.Println(attempts)
}
