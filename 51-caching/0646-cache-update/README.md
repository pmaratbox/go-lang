# 0646 — Update a value

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
It creates a cache with `lru.New[string,int](3)`, stores `Add("a", 1)`, then
re-puts the same key with `Add("a", 2)`. Adding an existing key overwrites its
value (rather than creating a duplicate), so `Get("a")` now returns the updated
value and prints `2`.

## Run

    go run .
