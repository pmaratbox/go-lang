# 0518 — Multiple arguments

Uses the [gopter](https://github.com/leanovate/gopter) property-testing library
run programmatically. `prop.ForAll` accepts multiple generators (`gen.Int()`,
`gen.Int()`), so the predicate receives two generated integer arguments `a` and
`b`. The property `max(a, b) >= a && max(a, b) >= b` holds for every pair, so
`Properties.Run` (fed a discarding reporter) returns `true` and we print
`passed`.

## Run

    go run .
