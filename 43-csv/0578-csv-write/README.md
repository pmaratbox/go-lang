# 0578 — Write CSV

Write rows to CSV with Go's standard-library `encoding/csv` package.
`csv.NewWriter` buffers a `bytes.Buffer`, and `WriteAll` encodes the two
rows `["name","age"]` then `["Alice","30"]` into CSV text. Go's writer emits
`\r\n` line terminators, so the output is normalized to `\n` and the trailing
newline is stripped before printing `name,age\nAlice,30`.

## Run

    go run .
