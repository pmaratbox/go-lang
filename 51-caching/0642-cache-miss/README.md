# 0642 — Cache miss

Uses the real strict-LRU cache library [hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
A fresh `lru.New[string, int](3)` cache is created and never populated, so the
lookup `Get("x")` returns `ok == false`. When a key is absent the program prints
`miss` instead of a value, demonstrating how the cache reports a missing entry.

## Run

    go run .
