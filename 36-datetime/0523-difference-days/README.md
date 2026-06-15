# 0523 — Difference in days

Use Go's standard `time` library to parse the two fixed ISO dates `2026-06-15` and `2026-07-15` with `time.Parse` (reference layout `2006-01-02`), then compute the gap between them with `Time.Sub`, converting the resulting `Duration` to whole days and printing `30`. No current-time call is used; both instants are fixed and the difference is computed by the library.

## Run

    go run .
