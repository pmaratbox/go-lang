# 0632 — Level filtering

Uses Go's standard-library structured logger `log/slog` with a `slog.NewJSONHandler` writing to an in-memory `bytes.Buffer` (no console, no real timestamp — the `ReplaceAttr` hook drops `slog.TimeKey`). The handler's minimum level is set to `slog.LevelWarn`, so the INFO record `hidden` is suppressed by the library and never reaches the buffer; only the WARN record `visible` is captured. That single captured JSON line is parsed back and re-printed as the normalized `level|message` form.

## Run

    go run .
