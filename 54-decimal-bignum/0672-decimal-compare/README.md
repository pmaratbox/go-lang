# 0672 — Exact decimal comparison

This lesson uses Go's `math/big.Rat` exact rational type to check whether `0.1 + 0.2` equals `0.3`. With binary floating point this is famously `false`, but because `big.Rat` stores values as exact fractions, the sum is exact and `Cmp` returns `0`, so the comparison prints `true`.

## Run

    go run .
