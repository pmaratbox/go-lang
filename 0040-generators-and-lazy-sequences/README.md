# 0040 — Generators & Lazy Sequences

Produce an endless lazy sequence of squares and take only the first three, printing `1 4 9`. Go (before the 1.23 range-over-func iterators) has no generators; a *stateful closure* is the idiom — `squares` returns a function that yields the next square each call. A goroutine writing to a channel is the other option.

## Run

    go run .
