# 0148 — Parse or Default

Parse "42" to 42 and "x" (invalid) to a default 0, printing `42 0`. Go's `strconv.Atoi` returns a value and an error, so you branch on `err` to fall back to a default.

## Run

    go run .
