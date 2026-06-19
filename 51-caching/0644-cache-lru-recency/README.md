# 0644 — Recency promotion

Uses the real strict-LRU cache library
[hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru).
With capacity 3 it stores `a=1, b=2, c=3`, then calls `Get("a")` which
promotes `a` to most-recently-used. Adding `d=4` overflows the cache, so the
now-least-recently-used key `b` is evicted instead of `a`. Looking up the
promoted `a` returns its value while the evicted `b` reports a miss, printing
`1 miss`.

## Run

    go run .
