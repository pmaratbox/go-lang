# 0516 — Sort is idempotent

Check a property over generated lists with the `github.com/leanovate/gopter` property-testing library, driven programmatically (no `testing.T`). `prop.ForAll` pairs the predicate with the `gen.SliceOf(gen.Int())` generator so gopter synthesizes ~100 random integer lists and verifies that sorting an already-sorted list equals sorting once (`sort(sort(xs)) == sort(xs)`). `gopter.NewProperties` runs with `DefaultTestParametersWithSeed(42)` and `Properties.Run` reports through a `NewFormatedReporter` writing to `io.Discard`, returning `true` when every case holds — so only `passed` reaches stdout.

## Run

    go run .
