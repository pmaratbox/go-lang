# 0320 — Fenwick Tree Prefix Sum

Build a Fenwick (BIT) tree over [1,2,3,4,5] and query the prefix sum of the first 4 elements, printing `10`. The `i & -i` low-bit trick drives both update and prefix loops, idiomatic Go for a BIT.

## Run

    go run .
