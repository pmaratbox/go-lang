# 0019 — Recursion

Define a recursive `factorial(n)` that multiplies `n` by `factorial(n - 1)` until it bottoms out at `1`, then print `factorial(5) = 120`. Go performs no tail-call optimization, but goroutine stacks start small and grow on demand, so recursion depth is limited by available memory rather than a fixed frame count. `int` is 64-bit on this platform.

## Run

    go run .
