# 0542 — Map / transform

Uses the `github.com/vektah/goparsify` parser-combinator library. The `Regex` combinator matches one-or-more digits, and the `.Map` combinator transforms the matched `Token` by parsing it with `strconv.Atoi` and doubling it, storing the result in `Result.Result`. Running the parser over the input `21` yields `21 * 2 = 42`.

## Run

    go run .
