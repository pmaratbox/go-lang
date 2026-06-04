# 0381 — Parallel Map

Square 1,2,3,4 in parallel and collect the results in input order, printing `1 4 9 16`. Each goroutine writes to its own index in a pre-sized slice, so order is preserved without locks.

## Run

    go run .
