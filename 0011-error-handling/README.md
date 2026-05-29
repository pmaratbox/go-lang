# 0011 — Error Handling

Write a `divide(a, b)` that reports a zero divisor, then call it on `10 / 2`
(prints the result) and `10 / 0` (prints an error). Go has no exceptions for
ordinary failures: a function returns an `error` as its last value, and callers
check `if err != nil`. `errors.New` builds a simple error.

## Run

    go run .
