# 0626 — Log at warn level

Uses Go's standard-library structured logger `log/slog` with a `slog.NewJSONHandler` writing to an in-memory `bytes.Buffer`. A `ReplaceAttr` callback strips the `time` attribute so no real timestamp appears. A single WARN record with the message `low disk` is emitted via `log.Warn`, then the captured JSON is parsed and re-emitted as one normalized line: the level (`WARN` mapped to `warn`) and message joined by `|`, followed by any structured fields sorted by key.

## Run

    go run .
