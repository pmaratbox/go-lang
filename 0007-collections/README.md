# 0007 — Collections

Build a slice of the integers `1, 2, 3, 4, 5`, then print its count and its
first and last elements. `[]int{...}` is a slice — a growable view over a
backing array — and `len()` is a builtin. Go has no negative indexing; the
last element is `nums[len(nums)-1]`. A fixed-size array (`[5]int`) is a
distinct type from a slice.

## Run

    go run .
