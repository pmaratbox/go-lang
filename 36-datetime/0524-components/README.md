# 0524 — Date components

Use Go's standard `time` library to parse the fixed ISO date `2026-06-15` with `time.Parse` (reference layout `2006-01-02`), then extract its components via the library's accessors `Year()`, `Month()`, and `Day()`, printing each on its own line: `2026`, `6`, `15`. The instant is fixed and every component is computed by the library, never from the current time.

## Run

    go run .
