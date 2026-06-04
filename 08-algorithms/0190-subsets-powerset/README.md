# 0190 — Subsets (Power Set)

Generate the power set of [1,2,3] in bitmask order 0..7 (empty printed as `{}`), one subset per line space-separated. Idiomatic Go iterates masks with `1<<n` and tests bits via `mask&(1<<i)`.

## Run

    go run .
