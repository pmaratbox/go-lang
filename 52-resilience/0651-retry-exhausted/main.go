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
		return errors.New("always fails") // never succeeds
	}
	// Allow up to 2 retries with zero delay => 3 total attempts. The op never
	// succeeds, so the library exhausts its retries and returns the last error.
	err := backoff.Retry(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 2))
	if err != nil {
		fmt.Println("failed")
	}
}
