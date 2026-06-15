# 0514 — Integer property

Uses the `github.com/leanovate/gopter` property-testing library run
programmatically: `prop.ForAll` with two `gen.Int()` generators checks that
addition is commutative (`a+b == b+a`) over ~100 generated integer pairs.
`Properties.Run` returns `true` when every case holds, so we print `passed`;
the reporter is wired to `io.Discard` to keep stdout clean.

## Run

    go run .
