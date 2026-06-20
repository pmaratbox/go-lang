# 0699 — Count added lines

Diff two fixed line-lists with the real diff library
[`github.com/pmezard/go-difflib/difflib`](https://github.com/pmezard/go-difflib),
a Go port of Python's `difflib`. `difflib.NewMatcher(a, b).GetOpCodes()` returns
an LCS-based edit script of opcodes; `'i'` (insert) and the target side of `'r'`
(replace) opcodes contribute ADDED lines. Diffing
`A=[apple,banana,cherry]` into `B=[apple,blueberry,cherry,date]` adds
`blueberry` and `date`, so the printed added count is `2`.

## Run

    go run .
