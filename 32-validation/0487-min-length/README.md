# 0487 — Minimum length

Validate a struct with Go's real `go-playground/validator/v10` library. The `Name` field carries a `validate:"min=3"` tag, so the string `"al"` (length 2) violates the minimum-length constraint. The program runs `v.Struct` on the input and prints the lowercased, sorted failing field name(s) pulled from the returned `validator.ValidationErrors` — here `name`.

## Run

    go run .
