package main

import (
	"errors"
	"fmt"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	attempts := 0
	// Scripted failure: fail once, then return the string "ok".
	op := func() (string, error) {
		attempts++
		if attempts < 2 {
			return "", errors.New("transient failure")
		}
		return "ok", nil
	}
	// RetryWithData drives the retries (up to 4 zero-delay retries) and
	// hands back the value produced on the successful attempt.
	result, err := backoff.RetryWithData(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 4))
	if err != nil {
		fmt.Println("failed")
		return
	}
	// Print the returned value, not the attempt count.
	fmt.Println(result)
}
