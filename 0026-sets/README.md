# 0026 — Sets

Build a set from `1, 2, 2, 3` so the duplicate collapses, then print its `size: 3` and whether it contains `2` (`has 2: yes`) and `5` (`has 5: no`). Go has no set type, so a `map[int]struct{}` is used as one — the empty struct holds no data, so the keys *are* the set. `len` gives the size, and the comma-ok form `_, ok := m[k]` tests membership.

## Run

    go run .
