# 0461 — Nested object

Uses Go's standard-library `encoding/json` package to serialize a nested
struct. A `Person` embeds an `Address` value; both structs declare their
fields alphabetically and carry `json:"..."` tags, so `json.Marshal`
produces compact JSON with keys already in alphabetical order, including the
nested object. The result is printed.

## Run

    go run .
