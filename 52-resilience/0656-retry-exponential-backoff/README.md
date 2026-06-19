# 0656 — Exponential backoff

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). It configures an
exponential backoff strategy via `backoff.NewExponentialBackOff()` with a zero
base interval, so the normally growing delays collapse to no waiting between
attempts. The scripted operation fails three times and then succeeds on its
fourth call, so `backoff.Retry` retries it three times and a closure counter
records four total attempts, printing `4`.

## Run

    go run .
