# 0671 — Exact decimal subtraction

This lesson uses Go's `math/big.Rat` exact rational type to subtract `1.0 - 0.1`. Because `big.Rat` stores values as exact fractions rather than binary floating point, the subtraction is exact and `FloatString(1)` prints `0.9` with no rounding error.

## Run

    go run .
