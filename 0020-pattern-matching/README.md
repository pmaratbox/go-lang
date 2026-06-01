# 0020 — Pattern Matching

Match `n` against the literal patterns `1` and `2` with a wildcard fallback, mapping `1`, `2`, and `5` to `one`, `two`, and `many`. Go's `switch` needs no `break` — each case breaks automatically (use `fallthrough` to chain). A conditionless `switch` acts like an if/else chain, and `default` handles the rest.

## Run

    go run .
