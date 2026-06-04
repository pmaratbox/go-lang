# 0328 — Trampoline

Sum 1..100 with a trampolined recursion that avoids deep stacks, printing `5050`. Each step returns a `thunk` closure that the driver loop invokes, keeping the stack flat.

## Run

    go run .
