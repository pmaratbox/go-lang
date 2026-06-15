# 0489 — Maximum length

Validate that a string does not exceed a maximum length using Go's real `github.com/go-playground/validator/v10` library. The `Code` field carries the `validate:"max=5"` tag, so the input `{code:'ABCDEFG'}` (7 characters) violates the `max` constraint. The program runs the validator, collects the failing field names from the `validator.ValidationErrors` result (lowercased and sorted), and prints them — here `code`.

## Run

    go run .
