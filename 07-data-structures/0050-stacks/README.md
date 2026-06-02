# 0050 — Stacks

Push `1`, `2`, and `3` onto a stack, then pop them all off and print them in last-in-first-out order: `3 2 1`. Go has no stack type; a slice is used directly — `append` pushes, and re-slicing off the last element (`s[:len(s)-1]`) pops. The last index is read before truncating.

## Run

    go run .
