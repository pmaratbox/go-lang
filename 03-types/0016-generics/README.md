# 0016 — Generics

Define a generic `first` function that returns the first element of a list, then call it on a list of integers and a list of strings to show one definition working at two types. Go added generics in 1.18: `func first[T any](items []T) T` declares a type parameter `T` constrained by `any` (the alias for `interface{}`). The compiler infers `T` from the argument at each call, so `first(ints)` and `first(strs)` need no explicit type argument. Constraints (here `any`, or interfaces like `comparable`) bound what a type parameter is allowed to do.

## Run

    go run .
