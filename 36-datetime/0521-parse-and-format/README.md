# 0521 — Parse and Format

Use Go's standard `time` library to parse the fixed ISO date `2026-06-15` with `time.Parse` (reference layout `2006-01-02`), then format the resulting value back to ISO `yyyy-MM-dd` with `Format`, printing `2026-06-15`. No current-time call is used; the instant is fixed and round-tripped by the library.

## Run

    go run .
