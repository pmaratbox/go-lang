# 0422 — Throttle (Virtual Time)

Implement throttle(window) (leading edge) on a virtual-time scheduler: emit a value, then suppress further values for `window` ticks. Go's `container/heap` backs the scheduler's (time, seq) min-heap priority queue.

## Run

    go run .
