# 0324 — Either Monad

Chain Either computations: a successful divide chain yields 2, and a divide-by-zero yields an error, printing `2 err`. A struct carrying either a Right int or a Left tag lets `bind` short-circuit on the first error.

## Run

    go run .
