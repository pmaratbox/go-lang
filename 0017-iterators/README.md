# 0017 — Iterators

Take the numbers 1 through 5, keep the even ones, double each, and add them up — a filter, then a map, then a reduce — printing the final sum. Go deliberately omits map/filter/reduce from its core and favors an explicit `for ... range` loop: the filter is an `if`, the map is the `n * 2` expression, and the reduce is the running `sum` accumulator. The result is more verbose but allocation-free and obvious in its control flow. (Later Go versions add range-over-function iterators, but the plain loop remains the idiom here.)

## Run

    go run .
