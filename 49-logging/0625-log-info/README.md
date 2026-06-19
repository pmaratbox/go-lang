# 0625 — Log at info level

Uses Go's standard-library structured logger `log/slog` with a `slog.NewJSONHandler` writing to an in-memory `bytes.Buffer` (no console, no real timestamp — the `ReplaceAttr` hook drops `slog.TimeKey`). An INFO record with the message `service started` and no fields is emitted, then the captured JSON line is parsed back and re-printed as the normalized `level|message` form.

## Run

    go run .
