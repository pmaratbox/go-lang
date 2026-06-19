# 0667 — Big integer power

Uses Go's arbitrary-precision integer type `big.Int` from `math/big`. The
`Exp` method computes `base**exponent` exactly, so `2^100` is produced without
any floating-point rounding and printed in full.

## Run

    go run .
