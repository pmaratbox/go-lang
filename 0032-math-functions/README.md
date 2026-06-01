# 0032 — Math Functions

Take the square root of `16`, raise `2` to the 10th power, the absolute value of `-5`, and the larger of `3` and `9`, printing `sqrt: 4`, `pow: 1024`, `abs: 5`, and `max: 9`. The `math` package provides `Sqrt`, `Pow`, and `Abs`, all operating on `float64` (cast to `int` here). The two-argument `max` (and `min`) are builtins added in Go 1.21, so no import is needed for those.

## Run

    go run .
