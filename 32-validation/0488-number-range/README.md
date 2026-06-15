# 0488 — Number range

Validate that a number falls within a range using Go's real `github.com/go-playground/validator/v10` library. The `Age` field carries the `validate:"min=0,max=120"` tag, so the input `{name:'alice', age:200}` violates the `max` constraint. The program runs the validator, collects the failing field names from the `validator.ValidationErrors` result (lowercased and sorted), and prints them — here `age`.

## Run

    go run .
