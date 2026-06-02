# 0047 — Immutable Update (Copy-with)

Make a copy of the point `(1, 2)` with its `x` changed to `9`, leaving the original intact, and print `original: (1, 2)` then `updated: (9, 2)`. Go structs are value types, so `p2 := p1` copies the whole struct; mutating `p2.X` cannot affect `p1`. There is no copy-with syntax because a plain assignment already copies.

## Run

    go run .
