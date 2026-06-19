# 0649 — Succeeds first try

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). It wraps an
operation in `backoff.Retry` with `WithMaxRetries(&backoff.ZeroBackOff{}, 4)`,
allowing up to four zero-delay retries. The operation returns `nil` on its very
first call, so the library never retries and a closure counter records a single
attempt, printing `1`.

## Run

    go run .
