# 0631 — Deterministic field order

Uses Go's standard-library structured logger `log/slog` with a `slog.NewJSONHandler` writing to an in-memory `bytes.Buffer` (no console, no real timestamp — the `ReplaceAttr` hook drops `slog.TimeKey`). An INFO record with the message `metric` is emitted with two fields supplied in non-alphabetical order (`zeta=2` then `alpha=1`); the captured JSON line is parsed back and re-printed with its fields sorted by key as `level|message|key=value...`.

## Run

    go run .
