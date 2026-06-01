# 0025 — Tuples & Multiple Return

Return both the smaller and larger of `3` and `7` from one function, unpacking the pair to print `min: 3` and `max: 7`. Multiple return values are first-class in Go: the signature `(int, int)` returns two results and `lo, hi := minMax(...)` receives them. This is the idiom behind the `value, err` error pattern; there is no tuple type.

## Run

    go run .
