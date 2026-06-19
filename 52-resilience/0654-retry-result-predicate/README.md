# 0654 — Retry on result

Uses the real retry library
[cenkalti/backoff/v4](https://github.com/cenkalti/backoff). The scripted
operation returns an incrementing counter, and the retry condition is based on
the returned VALUE rather than a thrown exception: the operation reports an
error (asking for another try) while the value is `< 3`, and succeeds once the
value reaches `3`. `backoff.RetryWithData` wraps it with
`WithMaxRetries(&backoff.ZeroBackOff{}, 4)`, driving up to four zero-delay
retries until the predicate is satisfied, then carries back the accepted
value. The program prints that accepted result, so the output is `3`.

## Run

    go run .
