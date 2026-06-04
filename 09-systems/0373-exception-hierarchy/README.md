# 0373 — Exception Hierarchy

Throw a specific error subtype and catch it through a base-type handler, printing `caught base`. In Go, errors.Is walks the Unwrap chain to match a wrapped base error.

## Run

    go run .
