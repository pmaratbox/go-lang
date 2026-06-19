# 0661 — Multiple dependencies

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig). Two services `A` (whose `X()`
returns `a`) and `B` (whose `Y()` returns `b`) are registered alongside a
`Service` whose constructor `NewService(*A, *B)` depends on BOTH. When
`Service` is resolved with `c.Invoke`, dig inspects the constructor's
parameters, builds the dependency graph, constructs `A` and `B`, injects
them, and hands back the wired `*Service`. Its `Run()` method returns
`X()+Y()`, printing `ab`.

## Run

    go run .
