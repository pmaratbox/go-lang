# 0115 — Dispatch Table

Store functions in a map keyed by name, then apply "add" and "mul" to (3,4), printing `7 12`. A Go `map[string]func(int, int) int` is the idiomatic dispatch table.

## Run

    go run .
