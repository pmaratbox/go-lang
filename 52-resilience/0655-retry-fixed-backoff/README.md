# 0655 — Fixed backoff

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). It configures a
fixed (constant) backoff strategy via `backoff.NewConstantBackOff(0)`, which
waits the same zero delay between every attempt. The scripted operation fails
twice and then succeeds on its third call, so `backoff.Retry` retries it twice
and a closure counter records three total attempts, printing `3`.

## Run

    go run .
