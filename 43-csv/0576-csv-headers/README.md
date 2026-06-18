# 0576 — CSV header row

Uses Go's standard `encoding/csv` package to parse a fixed CSV document. The
reader's `ReadAll` returns `[][]string`; the first row is the header. We join
the header fields with a pipe to produce `name|age|city`.

## Run

    go run .
