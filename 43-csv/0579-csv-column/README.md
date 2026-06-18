# 0579 — Extract a column

Uses Go's standard `encoding/csv` package to parse a fixed CSV document. The
reader's `ReadAll` returns `[][]string`. We locate the `age` column by its name
in the header row, pull that field from each data row, and join the values with
commas to produce `30,25,35`.

## Run

    go run .
