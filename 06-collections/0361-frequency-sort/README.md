# 0361 — Frequency Sort

Sort [1,1,2,3,3,3] by descending frequency (ties keep first-seen order), printing `3 3 3 1 1 2`. A first-seen `order` slice plus `sort.SliceStable` keeps ties stable in Go.

## Run

    go run .
