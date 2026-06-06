# 0405 — Map Operator

Implement a map operator that transforms each emitted value, applying x => x*2 to a stream of 1, 2, 3, 4. Go models the observer as a struct of next/error/complete closures, and the map operator wraps the source's subscribe.

## Run

    go run .
