# 0641 — Put and get

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
It creates a cache with `lru.New[string,int](3)`, stores a value with
`Add("a", 1)`, then retrieves it with `Get("a")` (which returns the value
and an `ok` flag and promotes the key to most-recently-used). The retrieved
value prints `1`.

## Run

    go run .
