# 0643 — LRU eviction

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
With `lru.New[string,int](3)` the cache holds at most three keys. Adding
`a=1, b=2, c=3, d=4` with no intervening `Get` calls overflows the
capacity, so the least-recently-used key `a` is evicted automatically.
Looking up `a` then returns `miss`, while `d` is still present and returns
`4`, printing `miss 4`.

## Run

    go run .
