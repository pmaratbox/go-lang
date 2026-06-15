# 0520 — Finds a counterexample

Uses the [gopter](https://github.com/leanovate/gopter) property-testing library
run programmatically. The property "every non-negative integer is `< 100`" is
densely false, so `prop.ForAll` over `gen.IntRange(0, 1000000)` generates an
input that violates it. `Properties.Run` returns `false`, which we detect to
print `found`. The reporter is fed `io.Discard`, so gopter's falsifying-example
and shrink report never leak to stdout — only `found` is printed.

## Run

    go run .
