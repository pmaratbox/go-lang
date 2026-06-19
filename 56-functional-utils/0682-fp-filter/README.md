# 0682 — Filter

Uses Go's `github.com/samber/lo` functional library and its `lo.Filter`
transform to keep only the elements of `[1,2,3,4,5,6]` that satisfy the
predicate `x%2 == 0` (the even numbers). The surviving values are mapped to
strings and comma-joined to print `2,4,6`.

## Run

    go run .
