# 0666 — Big integer multiplication

Multiplying two large integers with Go's arbitrary-precision `big.Int` from `math/big`. We build the operands with `big.NewInt` and compute the exact product via `(*big.Int).Mul`, so the result `123456789 * 987654321 = 121932631112635269` is exact with no overflow or float rounding.

## Run

    go run .
