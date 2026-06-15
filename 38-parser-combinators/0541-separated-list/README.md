# 0541 — Separated list

Parse a comma-separated list. Using the `github.com/vektah/goparsify` library, a `sepBy` parser is built from the `Seq`, `Many`, and `Map` combinators: a leading integer followed by `Many` repetitions of (`,` then integer). Running it on `"1,2,3"` yields `[1 2 3]`, whose elements sum to `6`.

## Run

    go run .
