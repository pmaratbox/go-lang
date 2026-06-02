# 0091 — Balanced Parentheses

Using a stack, check whether `(())` is balanced and whether `(()` is balanced, printing `yes` then `no`. A slice is the stack: append on `(`, reslice to drop the top on `)`; a non-empty stack at the end means unbalanced.

## Run

    go run .
