# 0665 — Big integer addition

Uses Go's stdlib `math/big.Int` arbitrary-precision integer type to add two
huge values that overflow any fixed-width integer. Each operand is parsed with
`SetString(s, 10)`, and `new(big.Int).Add(a, b)` computes the exact sum
`12345678901234567890 + 98765432109876543210`, printing
`111111111011111111100`.

## Run

    go run .
