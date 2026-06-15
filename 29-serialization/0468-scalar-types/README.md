# 0468 — Scalar types

Uses the standard library `encoding/json` package to marshal a struct with the
three scalar field types — `bool`, `int`, and `string` — into compact JSON.
Fields are declared alphabetically (`active`, `count`, `label`) so the
declaration-order serializer emits alphabetical keys, and Go renders the boolean
as lowercase `true`.

## Run

    go run .
