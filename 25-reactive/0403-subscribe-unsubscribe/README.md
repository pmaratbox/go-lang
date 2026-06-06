# 0403 — Subscribe and Unsubscribe

Return a Subscription from subscribe() and use it to unsubscribe so later values are not delivered. In Go a closure capturing a shared `closed` flag by pointer is the idiomatic way to wire the Subscription back into the producer's loop.

## Run

    go run .
