# 0660 — Bind interface to impl

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig). The `NewAnimal` constructor is
registered with `c.Provide`, but it returns the `Animal` interface rather than
the concrete `*Dog`, so dig records the implementation under the abstraction.
`c.Invoke(func(a Animal){...})` then resolves the service *by the interface*,
and the bound `*Dog`'s `Sound()` method returns and prints `woof`.

## Run

    go run .
