# 0086 — Sort a Map by Value

Sort the map `{a: 3, b: 1, c: 2}` by value in ascending order and print the entries: `b:1 c:2 a:3`. The keys are collected and `sort.Slice`-d by their mapped value, since Go maps have no order and cannot be sorted directly.

## Run

    go run .
