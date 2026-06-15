# 0491 — Multiple errors

Validate a struct with Go's real `go-playground/validator/v10` library. The `Name` field carries a `validate:"min=3"` tag and `Age` carries `validate:"min=0,max=120"`, so the input `{Name:"al", Age:200}` violates BOTH constraints at once. The validator reports every failing field (it does not stop at the first), and the program collects the lowercased, deduped, sorted failing field names from the returned `validator.ValidationErrors` — here `age` then `name`.

## Run

    go run .
