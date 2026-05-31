# 0018 — Closures

Build a counter that captures a private count starting at zero; each call to the returned function increments the count and returns it, so calling it twice prints 1 then 2. Go functions are first-class values and close over their enclosing variables by reference: the returned `func() int` keeps `count` alive on the heap after `counter` returns. Each returned closure owns its own `count`, so calling `counter()` again would start a fresh, independent counter.

## Run

    go run .
