# 0259 — Enumerate Submasks

Enumerate all submasks of the mask 5 (101) in descending order `5 4 1 0`. A bare `for` loop with the `(sub-1)&mask` step idiom walks every submask down to zero.

## Run

    go run .
