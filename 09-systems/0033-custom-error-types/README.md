# 0033 — Custom Error Types

Define a custom error, raise it from a `check` that rejects values over `100`, catch it for the input `200`, and print `error: value too large`. Go has no exceptions — errors are ordinary values. A type becomes an error by implementing `Error() string`; the function returns it, and the caller checks `err != nil`. `errors.Is`/`errors.As` compare or unwrap specific error types.

## Run

    go run .
