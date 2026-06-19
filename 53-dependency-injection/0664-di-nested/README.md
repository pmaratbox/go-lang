# 0664 — Nested dependency chain

Uses the real DI container
[uber/dig](https://github.com/uber-go/dig) to resolve a 3-level dependency
chain. Three constructors are registered with `c.Provide`: `A`, `B(*A)`, and
`C(*B)`. Resolving `C` with `c.Invoke` makes dig walk the graph in
order — constructing `A`, injecting it into `B`, then injecting `B` into `C` —
so that `C.V()` returns `A.V()+"b"+"c"`, printing `abc`.

## Run

    go run .
