# 0414 — FlatMap (mergeMap)

Implement flatMap/mergeMap: map each outer value to an inner timed stream and merge all inners concurrently (no cancellation). A `container/heap` min-heap ordered by (time, seq) drives the virtual-time scheduler deterministically.

## Run

    go run .
