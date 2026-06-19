package main

import (
	"fmt"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	attempts := 0
	op := func() error {
		attempts++
		return nil // succeeds on the first call
	}
	// Allow up to 4 retries with zero delay; the op succeeds immediately,
	// so the library never retries and the counter stays at 1.
	backoff.Retry(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 4))
	fmt.Println(attempts)
}
