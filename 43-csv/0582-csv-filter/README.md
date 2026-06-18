# 0582 — Filter rows

Parse a `name,age,city` CSV with Go's standard `encoding/csv` library and filter
its data rows by a predicate. `csv.NewReader(...).ReadAll()` decodes the fixed
document into a `[][]string`; we skip the header, keep each row whose `age` is
greater than 28 (Alice 30 and Carol 35 qualify, Bob 25 does not), and print the
kept names comma-joined as `Alice,Carol`.

## Run

    go run .
