# 0081 — Set Operations

Compute the union and intersection of the sets `{1, 2, 3}` and `{2, 3, 4}`, printing the union `1 2 3 4` and the common elements `2 3` (each in ascending order). Go models sets as `map[int]bool`; union merges both maps and intersection keeps `a`'s keys present in `b`, then the keys are sorted.

## Run

    go run .
