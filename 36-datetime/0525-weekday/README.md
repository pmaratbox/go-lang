# 0525 — Weekday

Using go's stdlib `time` library, parse the fixed date `2026-06-15` and read its `Weekday()`. Go's `time.Weekday` is Sunday-based (Sunday=0), so we convert it to the ISO weekday number (Monday=1 .. Sunday=7) with `(int(d.Weekday())+6)%7 + 1`, yielding `1` for this Monday.

## Run

    go run .
