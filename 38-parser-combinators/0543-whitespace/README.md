# 0543 — Whitespace handling

Uses the `github.com/vektah/goparsify` parser-combinator library. The `Regex` combinator matches one-or-more digits and `.Map` folds the matched `Token` into an `int`. Because goparsify's combinators auto-skip surrounding whitespace (leading space is consumed before each token and `Run` discards trailing space), parsing the input `  42  ` cleanly yields the integer `42` without any manual trimming.

## Run

    go run .
