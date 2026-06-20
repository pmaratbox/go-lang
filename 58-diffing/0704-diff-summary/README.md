# 0704 — Diff summary

Diff two fixed line-lists with the real diff library
[`github.com/pmezard/go-difflib/difflib`](https://github.com/pmezard/go-difflib),
a Go port of Python's `difflib`. `difflib.NewMatcher(a, b).GetOpCodes()` returns
an LCS-based edit script of opcodes: `'i'` (insert) contributes added lines,
`'d'` (delete) contributes removed lines, `'r'` (replace) maps to both
removed(source) and added(target), and `'e'` (equal) counts the unchanged lines.
Diffing `A=[apple,banana,cherry]` into `B=[apple,blueberry,cherry,date]` yields
2 added, 1 removed, and 2 unchanged, printed space-joined as `2 1 2`.

## Run

    go run .
