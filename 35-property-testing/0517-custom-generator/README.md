# 0517 — Custom generator

Uses the gopter property-testing library run programmatically (no `testing.T`). A custom generator is built with the `Map` combinator — `gen.Int().Map(n -> n*2)` — so every produced value is even; the property asserts each generated value satisfies `n%2 == 0`. `Properties.Run` with a discarding reporter keeps stdout clean, printing only `passed`.

## Run

    go run .
