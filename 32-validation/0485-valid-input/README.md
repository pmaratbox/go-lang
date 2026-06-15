# 0485 — Valid input

Validate a struct with Go's real `github.com/go-playground/validator/v10` library. The schema constrains `Name` with `min=3` and `Age` with `min=0,max=120` via `validate` struct tags. The valid input `{name: alice, age: 30}` satisfies every constraint, so `validator.Struct` returns a nil error and the program prints `ok`. On failure it would instead print the sorted, lowercased failing field name(s) drawn from the validator's `ValidationErrors`.

## Run

    go run .
