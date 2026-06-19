# 0690 — Flat map

This lesson uses the `github.com/samber/lo` functional library and its `lo.FlatMap` transform to map each element of `[1,2,3]` to the slice `[x, x*10]` and flatten the results into one slice. We then comma-join the values to print `1,10,2,20,3,30`.

## Run

    go run .
