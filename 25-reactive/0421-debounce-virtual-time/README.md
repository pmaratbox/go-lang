# 0421 — Debounce (Virtual Time)

Implement debounce(window) on a virtual-time scheduler: emit a value only after a quiet gap of `window` ticks with no newer value. In Go a `container/heap` min-heap ordered by (time, seq) drives the scheduler, and observers are plain structs holding `next`/`complete` closures.

## Run

    go run .
