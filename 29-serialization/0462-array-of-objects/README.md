# 0462 — Array of objects

Uses the standard library `encoding/json` package and `json.Marshal` to
serialize a Go slice of structs into a single compact JSON array. Each
`Person` struct's fields are declared alphabetically and carry `json:"..."`
tags, so the marshaled object keys come out compact and in alphabetical
order (`age`, `name`).

## Run

    go run .
