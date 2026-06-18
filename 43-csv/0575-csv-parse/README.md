# 0575 — Parse CSV rows

Parse CSV text with Go's standard-library `encoding/csv` package.
`csv.NewReader(strings.NewReader(data)).ReadAll()` decodes the fixed CSV
`name,age,city\nAlice,30,Paris\n...` into a `[][]string`. The header row is
skipped, the first column (name) of each data row is collected, and the
values are joined with commas via `strings.Join` to print `Alice,Bob,Carol`.

## Run

    go run .
