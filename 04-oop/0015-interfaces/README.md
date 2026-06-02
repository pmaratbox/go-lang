# 0015 — Interfaces

Define a `Shape` interface with `name()` and `area()` methods, implement it for a rectangle and a square, then loop over a collection of shapes and print each one's area. Go interfaces are *structural* and implicit: there is no `implements` keyword — `Rectangle` and `Square` satisfy `Shape` simply by declaring the right methods. A `[]Shape` holds either concrete type, and ranging over it calls `s.area()` through an interface value, which dispatches to the underlying type at runtime.

## Run

    go run .
