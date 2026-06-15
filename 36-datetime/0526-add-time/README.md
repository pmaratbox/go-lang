# 0526 — Add time

Using Go's standard `time` library, we parse the fixed instant
`2026-06-15T10:00`, add 90 minutes with `time.Time.Add(90 * time.Minute)`, and
format the result as `HH:mm` with the reference layout `15:04`. The arithmetic
is done by the library, not hardcoded.

## Run

    go run .
