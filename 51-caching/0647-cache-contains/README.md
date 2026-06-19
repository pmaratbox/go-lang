# 0647 — Contains key

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
It creates a cache with `lru.New[string,int](3)` and stores one entry with
`Add("a", 1)`. The `Contains` method reports key membership *without*
promoting the key to most-recently-used (unlike `Get`). Checking `a` returns
`true` and checking the absent `x` returns `false`, printed lowercase and
space-joined as `true false`.

## Run

    go run .
