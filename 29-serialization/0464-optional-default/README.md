# 0464 — Optional field with default

Uses the standard library `encoding/json` package's `json.Unmarshal` to parse
an object where the `age` field is absent. Because Go decodes JSON into a
struct, any field missing from the input keeps its zero value, so `Age` defaults
to `0` without any extra handling.

## Run

    go run .
