# 0539 — Alternative

Uses the `github.com/vektah/goparsify` parser-combinator library. The `Any` combinator expresses an alternative (choice): it tries the literal `cat` first and falls back to the literal `dog`. Running the parser over the input `dog` matches the second branch and yields `dog`.

## Run

    go run .
