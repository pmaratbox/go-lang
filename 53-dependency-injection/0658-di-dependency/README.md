# 0658 — Inject a dependency

Uses the real DI container [uber-go/dig](https://github.com/uber-go/dig). Two
constructors are registered with `Provide`: `NewRepo` builds a `*Repo` whose
`Data()` returns `data`, and `NewService` declares a `*Repo` parameter so dig
injects the resolved `Repo` into the `Service`. Resolving `*Service` via
`Invoke` walks the dependency graph, and calling `Run()` delegates to the
injected repo, printing `data`.

## Run

    go run .
