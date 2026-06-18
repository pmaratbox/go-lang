# 0577 — Quoted CSV fields

Parse a quoted CSV field with Go's standard-library `encoding/csv` package.
`csv.NewReader(strings.NewReader(data)).ReadAll()` decodes
`name,note\nAlice,"hello, world"\n` into a `[][]string`. The `note` value is
wrapped in double quotes so its embedded comma is part of the field rather than
a column separator; the reader handles the quoting and yields the intact value
`hello, world`, which is printed from the data row.

## Run

    go run .
