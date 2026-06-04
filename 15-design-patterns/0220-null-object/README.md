# 0220 — Null Object

Compare a no-op null logger with a real logger; only the real one records, so print the logged count `1`. The null object satisfies the same interface with empty method bodies, avoiding nil checks.

## Run

    go run .
