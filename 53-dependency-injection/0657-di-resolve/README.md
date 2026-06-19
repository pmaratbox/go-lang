# 0657 — Resolve a service

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig). A `Greeter` service is registered
with `c.Provide(NewGreeter)`, then resolved out of the container with
`c.Invoke`, which inspects the function's parameters, builds the dependency
graph, and supplies the constructed `*Greeter`. The resolved service's
`Greet()` method returns and prints `hello`.

## Run

    go run .
