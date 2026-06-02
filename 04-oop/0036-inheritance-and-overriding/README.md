# 0036 — Inheritance & Overriding

Define a base `Animal` with a `speak` method, a `Dog` that overrides it, and call both, printing `animal: some sound` and `dog: Woof`. Go has no inheritance; `Dog` *embeds* `Animal`, which promotes its methods, and declaring `Dog`'s own `Speak` shadows the embedded one. There is no virtual dispatch — a method on `Animal` would call `Animal.Speak`, not the override.

## Run

    go run .
