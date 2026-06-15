# 0519 — Precondition / filter

Constrain generated inputs with a precondition using the `github.com/leanovate/gopter` property-testing library, driven programmatically (no `testing.T`). `gen.Int().SuchThat(...)` filters the integer generator so only positive values reach the predicate, and `prop.ForAll` then verifies that `n + 1 > n` for every generated positive `n`. `gopter.NewProperties` runs with `DefaultTestParametersWithSeed(42)` and `Properties.Run` reports through a `NewFormatedReporter` writing to `io.Discard`, returning `true` when every case holds — so only `passed` reaches stdout.

## Run

    go run .
