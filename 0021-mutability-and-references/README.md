# 0021 — Mutability & References

Have a function increment a value in place — through a pointer, reference, or mutable holder — so the caller sees it change from `before: 1` to `after: 2`. Go passes arguments by value, so to mutate the caller's variable the function takes a pointer (`*int`) and the caller passes the address with `&n`. Inside, `*p++` increments the pointed-to value (pointer arithmetic like `p++` does not exist in Go). Slices and maps are reference-like and would mutate without an explicit pointer.

## Run

    go run .
