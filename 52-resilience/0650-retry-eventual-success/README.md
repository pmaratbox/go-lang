# 0650 — Eventual success

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). It wraps an
operation in `backoff.Retry` with `WithMaxRetries(&backoff.ZeroBackOff{}, 4)`,
allowing up to four zero-delay retries. The scripted operation returns an error
on its first call and `nil` on the second, so the library retries exactly once.
A closure counter records the total attempts driven by the library, printing
`2`.

## Run

    go run .
