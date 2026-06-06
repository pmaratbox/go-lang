# 0408 — Take Operator

Implement take(n) over an unbounded source of the natural numbers, emitting the first 3 then completing (and unsubscribing the source). The source loops while a `stopped` flag is false, and the take operator flips it via the returned Subscription to halt the otherwise infinite producer.

## Run

    go run .
