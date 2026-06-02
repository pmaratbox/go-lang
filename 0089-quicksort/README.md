# 0089 — Quicksort

Sort the list `3, 1, 4, 1, 5, 2` using quicksort (partition around a pivot, then recurse on each side) and print the result: `1 1 2 3 4 5`. One pass routes the rest into `less` or `greater`; the recursively sorted slices are appended around the pivot.

## Run

    go run .
