# 0486 — Required field

Validate a struct with Go's real `github.com/go-playground/validator/v10` library, where the `Age` field carries the `required` constraint. The input has `Name` present but `Age` missing (its zero value), so `v.Struct` returns a `validator.ValidationErrors`; the output is the failing field name(s) extracted via `e.Field()`, lowercased and sorted — here `age`.

## Run

    go run .
