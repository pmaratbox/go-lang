# 0515 — String property

Use the `github.com/leanovate/gopter` property-testing library to check a property over generated strings, run programmatically (not via `go test`). A `gopter.Properties` is built with a fixed seed and given `prop.ForAll(pred, gen.AlphaString())`: the `gen.AlphaString()` generator produces ~100 random strings `s`, and the predicate asserts `len(s+s) == 2*len(s)`. `props.Run` is driven with a `NewFormatedReporter` writing to `io.Discard`, so nothing leaks to stdout; it returns `true` when every generated case holds, and we print `passed`.

## Run

    go run .
