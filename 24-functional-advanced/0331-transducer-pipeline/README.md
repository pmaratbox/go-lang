# 0331 — Transducer Pipeline

Compose map(+1) with filter(even) and run it over [1,2,3,4], printing `2 4`. Transducers wrap a reducer so the whole pipeline runs in a single pass with no intermediate slices.

## Run

    go run .
