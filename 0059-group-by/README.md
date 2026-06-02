# 0059 — Group By

Group the words `one`, `two`, `three` by their length and print each length with its words, in ascending order of length: `3:[one,two] 5:[three]`. A `map[int][]string` accumulates each bucket via `append`. Map iteration is randomized, so the integer keys are collected and `sort.Ints`-ed before printing.

## Run

    go run .
