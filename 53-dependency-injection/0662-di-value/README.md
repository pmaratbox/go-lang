# 0662 — Inject a value

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig). A constant value `Cfg{V: "v1"}` is
registered with `c.Provide(NewCfg)`, and a `Service` whose constructor depends
on `Cfg` is registered with `c.Provide(NewService)`. Resolving the service with
`c.Invoke` builds the dependency graph, injects the constant value into the
service, and calling `Value()` returns and prints `v1`.

## Run

    go run .
