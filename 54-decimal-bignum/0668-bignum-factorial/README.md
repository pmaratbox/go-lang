# 0668 — Factorial

This lesson computes `30!` exactly using Go's arbitrary-precision integer type `math/big.Int`. Starting from 1, it repeatedly multiplies (`Mul`) by each integer from 1 through 30, so the result never overflows and carries every digit exactly — no floating-point rounding.

## Run

    go run .
