# 0702 — Removed line content

Diff two fixed line-lists with the real diff library
[`github.com/pmezard/go-difflib/difflib`](https://github.com/pmezard/go-difflib),
a Go port of Python's `difflib`. `difflib.NewMatcher(a, b).GetOpCodes()` returns
an LCS-based edit script of opcodes; `'d'` (delete) and the source side of `'r'`
(replace) opcodes contribute REMOVED lines. Diffing
`A=[apple,banana,cherry]` into `B=[apple,blueberry,cherry,date]` removes
`banana`, so the printed removed content is `banana`.

## Run

    go run .
