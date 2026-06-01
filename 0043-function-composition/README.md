# 0043 — Function Composition

Compose `inc` (add one) and `twice` (multiply by two) into one function and apply it to `3`, so `inc(twice(3))` prints `7`. `compose` takes two `func(int) int` values and returns a new closure `func(x) { return f(g(x)) }`. Functions are first-class in Go, so closures capture and carry the surrounding variables.

## Run

    go run .
