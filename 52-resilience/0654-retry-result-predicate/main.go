package main

import (
	"errors"
	"fmt"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	counter := 0
	// Scripted result sequence: the operation returns an incrementing
	// counter (1, 2, 3, ...). This is a retry-ON-RESULT condition, not a
	// retry-on-exception: we keep retrying while the produced VALUE is < 3.
	op := func() (int, error) {
		counter++
		value := counter
		if value < 3 {
			// Signal "not acceptable yet" so the library retries.
			return value, errors.New("result not acceptable yet")
		}
		// Acceptable result -> stop retrying and carry the value back.
		return value, nil
	}
	// RetryWithData drives the retries (up to 4 zero-delay retries) and
	// returns the value produced on the accepted attempt.
	result, err := backoff.RetryWithData(op, backoff.WithMaxRetries(&backoff.ZeroBackOff{}, 4))
	if err != nil {
		fmt.Println("failed")
		return
	}
	// Print the accepted result (the value that satisfied the predicate).
	fmt.Println(result)
}
