# 0195 — Defer LIFO Order

Register three deferred actions printing 1, 2, 3 and show they run in last-in-first-out order `3 2 1`. Go's `defer` stack runs deferred calls in last-in-first-out order at function return.

## Run

    go run .
