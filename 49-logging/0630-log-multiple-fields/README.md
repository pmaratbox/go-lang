# 0630 — Log multiple fields

Uses Go's standard-library structured logger `log/slog`. A `slog.JSONHandler`
writes records into an in-memory `bytes.Buffer` (with a `ReplaceAttr` that drops
the time key, so there is no real timestamp). We emit an INFO record `request`
with a string field `method=GET` and an integer field `status=200`, then parse
the captured JSON and print one normalized line `level|message` followed by each
field sorted by key as `|key=value` (JSON numbers come back as `float64` and are
rendered as plain integers).

## Run

    go run .
