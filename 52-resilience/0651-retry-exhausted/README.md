# 0651 — Retries exhausted

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). The operation
always returns an error, and `backoff.Retry` is wrapped with
`WithMaxRetries(&backoff.ZeroBackOff{}, 2)`, allowing up to 2 zero-delay
retries for 3 total attempts. Because every attempt fails, the library
exhausts its retries and returns the last error, which the program catches
and prints as `failed`.

## Run

    go run .
