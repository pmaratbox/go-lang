# 0627 — Log at error level

Uses Go's standard-library structured logger `log/slog`. An ERROR record with
the message `boom` is emitted through a `slog.JSONHandler` writing into an
in-memory `bytes.Buffer` (a `ReplaceAttr` hook strips the timestamp). The
captured JSON is parsed and reprinted as the normalized line `level|message`,
with the `ERROR` level mapped to the short lowercase form `error`.

## Run

    go run .
