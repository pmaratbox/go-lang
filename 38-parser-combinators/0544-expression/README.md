# 0544 — Expression

Uses the `github.com/vektah/goparsify` parser-combinator library. A `Regex` combinator parses each integer, and the `+`-separated sequence is built with `Seq` and `Many`: `Seq("+", num)` matches each trailing `+ integer` term, `Many` repeats it, and a final `.Map` folds all parsed integers into their sum. Running the parser over the input `10+20+30` yields `60`.

## Run

    go run .
