# 0412 — Combine Latest

Implement combineLatest of two timed streams, emitting the pair of latest values whenever either source emits (once both have emitted). A `container/heap` min-heap keyed by (time, seq) drives the virtual-time scheduler deterministically.

## Run

    go run .
