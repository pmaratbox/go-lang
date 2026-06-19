# 0659 — Singleton lifetime

Uses the real DI container [uber/dig](https://github.com/uber-go/dig). A `*Repo`
constructor is registered with `c.Provide(NewRepo)`, then the dependency is
resolved twice via separate `c.Invoke` calls. dig caches the value produced by a
constructor, giving every consumer the same instance, so the two resolved
pointers are identical and the identity check prints `true`.

## Run

    go run .
