# 0035 — Ranges & Slicing

From the list `[10, 20, 30, 40, 50]`, take the sub-sequence at indices 1 through 4 (exclusive) and print `slice: 20 30 40`. A slice expression `nums[1:4]` is half-open and returns a *view* that shares the underlying array — no copy — so writes through it affect the original. `len`/`cap` describe the view; `copy` makes an independent slice.

## Run

    go run .
