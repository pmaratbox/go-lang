# 0359 — Flatten Deeply

Flatten the arbitrarily nested structure [1,[2,[3,4]],5] into `1 2 3 4 5`. Go models the heterogeneous nesting with `[]any` and a type switch in the recursion.

## Run

    go run .
