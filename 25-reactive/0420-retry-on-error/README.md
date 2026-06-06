# 0420 — Retry On Error

Implement retry(n) that resubscribes to the source on error up to n times; the source succeeds on the 3rd subscription. A recursive closure capturing the remaining count drives the resubscription idiomatically in Go.

## Run

    go run .
