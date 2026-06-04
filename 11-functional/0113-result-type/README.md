# 0113 — Result / Either Type

Model success and failure with a Result type: safeDiv(10,2) prints `ok: 5` and safeDiv(1,0) prints `err: divide by zero`. A Go struct carrying value or error message mirrors an Either, matched in its Stringer.

## Run

    go run .
