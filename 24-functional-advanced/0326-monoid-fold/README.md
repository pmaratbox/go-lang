# 0326 — Monoid Fold

Fold lists under two monoids: string concat ["a","b","c"]->"abc" and integer sum [1,2,3]->6, printing `abc 6`. A single generic `fold` takes the identity and combine op for whichever monoid you supply.

## Run

    go run .
