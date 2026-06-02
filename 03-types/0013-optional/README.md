# 0013 — Optional

Hold one value that is present (`42`) and one that is absent, then print each
with a fallback of `-1` when absent. Go has no Option type — a pointer (`*int`)
models optionality, where `nil` means absent. The helper dereferences it or
returns the fallback.

## Run

    go run .
