# 0104 — Atomic Counter

Increment a shared atomic counter from multiple threads 1000 times total without a lock, printing `1000`. An atomic.Int64 with Add provides lock-free increments.

## Run

    go run .
