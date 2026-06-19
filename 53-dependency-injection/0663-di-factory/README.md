# 0663 — Factory provider

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig). Instead of letting dig autowire a
struct, the `Service` is registered through a FACTORY function,
`NewServiceFactory`, which explicitly constructs and configures the object. The
factory is handed to the container with `c.Provide`, and `c.Invoke` resolves the
factory-built `*Service` and calls its `Value()` method, which returns and prints
`built`.

## Run

    go run .
