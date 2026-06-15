# 0492 — Custom rule

Validate a struct with Go's real `go-playground/validator/v10` library using a custom rule. The rule `hasdigit` is registered via `v.RegisterValidation` and requires the value to contain at least one digit; the `Password` field carries the `validate:"hasdigit"` tag. The input `{Password:"abcdef"}` has no digit, so the rule fails, and the program prints the lowercased, sorted failing field name pulled from the returned `validator.ValidationErrors` — here `password`.

## Run

    go run .
