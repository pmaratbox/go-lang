# 0669 — Exact decimal addition

Adding `0.1 + 0.2` with binary floating point yields `0.30000000000000004`. Using Go's `math/big.Rat` (an arbitrary-precision exact rational type) we parse `0.1` and `0.2` exactly, `Add` them, and render the result with `FloatString(1)` to print the exact `0.3`.

## Run

    go run .
