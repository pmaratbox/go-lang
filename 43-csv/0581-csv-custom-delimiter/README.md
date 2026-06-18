# 0581 — Custom delimiter

Parse semicolon-delimited data with the `encoding/csv` standard-library
package. A `csv.Reader` is created over `a;b;c\n1;2;3\n`, and its `Comma`
field is set to `';'` so the parser splits on semicolons instead of commas.
`ReadAll` returns `[][]string`; the second (data) row's fields are joined
with commas via `strings.Join` to print `1,2,3`.

## Run

    go run .
