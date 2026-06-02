# 0066 — Multiple Assignment & Destructuring

Swap two variables (`a = 1`, `b = 2`) with a single multiple-assignment, then unpack the pair `(3, 4)` into two variables — printing `2 1` then `3 4`. Go's multiple assignment `a, b = b, a` swaps simultaneously (the right side is fully evaluated first); it has no array/tuple destructuring, so a pair is read element by element.

## Run

    go run .
