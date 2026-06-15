# 0527 — Timezone offset

Use Go's standard `time` library to parse the fixed UTC instant `2026-06-15T12:00:00Z` with `time.Parse` (RFC 3339 layout `2006-01-02T15:04:05Z07:00`), then convert it to a fixed `+05:00` offset zone built with `time.FixedZone` (no named timezone or OS tzdata) via `Time.In`. The local hour `17` is computed by the library, not hardcoded; no current-time call is used.

## Run

    go run .
