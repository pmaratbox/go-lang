# 0580 — Sum a numeric column

Sum a numeric CSV column with Go's standard-library `encoding/csv` package.
`csv.NewReader(strings.NewReader(data)).ReadAll()` decodes the fixed CSV
`name,age,city\nAlice,30,Paris\n...` into a `[][]string`. The header row is
skipped, the `age` column of each data row is parsed to an integer with
`strconv.Atoi`, and the values are accumulated (30+25+35) to print `90`.

## Run

    go run .
