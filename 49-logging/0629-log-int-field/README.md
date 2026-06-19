# 0629 — Log an integer field

Uses Go's standard-library structured logger `log/slog`. An INFO record with the
message `processed` carries one integer structured field `count=5`, emitted
through a `slog.JSONHandler` writing into an in-memory `bytes.Buffer` (a
`ReplaceAttr` hook strips the timestamp). The captured JSON is parsed — where
numbers arrive as `float64` and are formatted back without a decimal — and
reprinted as the normalized line `level|message|key=value`.

## Run

    go run .
