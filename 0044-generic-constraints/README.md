# 0044 — Generic Constraints

Write a generic `largest(a, b)` that requires an ordered type, then call it on integers (3 and 9) and on strings (apple and pear), printing `9` and `pear`. The type parameter `[T cmp.Ordered]` constrains `T` to types that support the ordering operators; `cmp.Ordered` (Go 1.21) is a union of all such built-in types, so `a > b` is allowed.

## Run

    go run .
