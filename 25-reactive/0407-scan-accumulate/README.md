# 0407 — Scan (Running Fold)

Implement a scan operator that emits the running accumulation; produce the running sums of 1, 2, 3, 4. A closure captures the accumulator state across pushed values, the idiomatic Go way to thread mutable state through callbacks.

## Run

    go run .
