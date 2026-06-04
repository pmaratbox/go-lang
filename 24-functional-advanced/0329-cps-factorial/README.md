# 0329 — CPS Factorial

Compute 5! in continuation-passing style, printing `120`. Each recursive call threads a continuation closure, and the top-level identity continuation returns the final product.

## Run

    go run .
