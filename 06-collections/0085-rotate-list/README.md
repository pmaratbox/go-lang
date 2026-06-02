# 0085 — Rotate a List

Rotate the list `1, 2, 3, 4, 5` left by `2` positions (elements wrap to the end) and print it: `3 4 5 1 2`. Appending the tail `nums[k:]` then the head `nums[:k]` into a fresh slice rotates left without aliasing the original.

## Run

    go run .
