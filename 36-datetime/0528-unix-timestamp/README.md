# 0528 — Unix timestamp

Compute the Unix timestamp (epoch seconds) of a fixed UTC instant with Go's standard `time` library. `time.Parse` reads the RFC 3339 string `2026-06-15T00:00:00Z` using the reference layout `2006-01-02T15:04:05Z07:00`, yielding a `time.Time` anchored to UTC (the trailing `Z` is a fixed zero offset, so no OS timezone database is consulted). Calling `.Unix()` returns the number of seconds elapsed since the Unix epoch (1970-01-01T00:00:00Z) — the value is derived by the library, not hardcoded.

## Run

    go run .
