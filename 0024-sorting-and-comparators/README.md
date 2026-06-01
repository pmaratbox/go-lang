# 0024 — Sorting & Comparators

Sort `[3, 1, 2]` ascending, then again with a custom comparator that reverses the order, printing `asc: 1 2 3` and `desc: 3 2 1`. `slices.Sort` orders a slice in place using natural ordering; `slices.SortFunc` takes a comparator returning negative/zero/positive, and `cmp.Compare(b, a)` reverses it. Both are generic helpers added in Go 1.21.

## Run

    go run .
