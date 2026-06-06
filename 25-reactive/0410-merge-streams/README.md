# 0410 — Merge Streams

Implement merge of two timed streams using a virtual-time scheduler, interleaving them by emission time. Go's container/heap backs the scheduler's (time, seq) min-heap for deterministic ordering.

## Run

    go run .
