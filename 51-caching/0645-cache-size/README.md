# 0645 — Cache size

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
It creates a cache with capacity 5 via `lru.New[string,int](5)`, stores two
entries with `Add("a", 1)` and `Add("b", 2)`, then reports how many entries
the cache currently holds with `Len()`. Since both keys fit under the
capacity, nothing is evicted and the size prints `2`.

## Run

    go run .
