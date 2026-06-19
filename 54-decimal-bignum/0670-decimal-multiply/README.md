# 0670 — Exact decimal multiplication

Multiply `1.1 * 1.1` using Go's `math/big.Rat`, an exact rational type, so the product is computed without floating-point rounding. The result is rendered with `FloatString(2)` to two decimal places, yielding exactly `1.21`.

## Run

    go run .
