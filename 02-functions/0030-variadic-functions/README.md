# 0030 — Variadic Functions

Define a function that accepts a variable number of integer arguments and returns their total, then call it with `1, 2, 3` to print `sum: 6`. A `...int` parameter makes the function variadic; inside, `nums` is an `[]int` slice ranged over to accumulate the sum. An existing slice is forwarded with `total(xs...)`.

## Run

    go run .
