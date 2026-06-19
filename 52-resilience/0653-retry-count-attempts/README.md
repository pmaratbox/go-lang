# 0653 — Count attempts

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). It wraps an
operation in `backoff.Retry` with `WithMaxRetries(&backoff.ZeroBackOff{}, 4)`,
allowing up to four zero-delay retries on top of the initial call. The
operation always returns an error, so the library makes every one of the five
total attempts before giving up; a closure counter records each call and prints
`5`.

## Run

    go run .
