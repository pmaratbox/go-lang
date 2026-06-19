# 0648 — Capacity bound

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
It creates a cache with `lru.New[string,int](3)`, then `Add`s four items
`a`, `b`, `c`, `d`. Because the cache is bounded to a capacity of 3, adding
the fourth item evicts the least-recently-used one, so the size never
exceeds the capacity. `Len()` reports `3`.

## Run

    go run .
