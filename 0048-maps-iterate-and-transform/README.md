# 0048 — Maps: Iterate & Transform

Build a map from letters to numbers (`a`->1, `b`->2, `c`->3), sum all its values, and print `sum: 6`. A `for range` over a map yields each key and value; here the value is accumulated. Map iteration order is randomized in Go, but summing is order-independent.

## Run

    go run .
