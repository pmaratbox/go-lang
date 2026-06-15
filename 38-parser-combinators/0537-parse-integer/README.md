# 0537 — Parse an integer

Uses the `github.com/vektah/goparsify` parser-combinator library. The `Regex` combinator matches one-or-more digits, and `.Map` folds the matched `Token` into an `int` via `strconv.Atoi`, storing it in `Result.Result`. Running the parser over the input `42` yields the integer `42`.

## Run

    go run .
