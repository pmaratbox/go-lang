# 0513 — First property

Check a property over generated lists with the `github.com/leanovate/gopter` property-testing library, driven programmatically (no `testing.T`). `prop.ForAll` pairs a predicate with the `gen.SliceOf(gen.Int())` generator so gopter synthesizes ~100 random integer lists and verifies that reversing a list twice yields the original. `gopter.NewProperties` runs with `DefaultTestParametersWithSeed(42)` and `Properties.Run` reports through a `NewFormatedReporter` writing to `io.Discard`, returning `true` when every case holds — so only `passed` reaches stdout.

## Run

    go run .
