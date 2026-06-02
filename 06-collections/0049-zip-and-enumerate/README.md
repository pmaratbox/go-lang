# 0049 — Zip & Enumerate

Pair the letters `a, b, c` with the numbers `1, 2, 3` position by position, formatting each pair as `key=value` and printing `a=1 b=2 c=3`. Go has no zip; a `for range` over one slice gives the index, used to reach into the parallel slice. The index-based loop is the standard pattern for walking two slices together.

## Run

    go run .
