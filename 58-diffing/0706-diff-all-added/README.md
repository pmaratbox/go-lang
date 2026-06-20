# 0706 — All lines added

Diff two fixed line-lists with the real diff library
[`github.com/pmezard/go-difflib/difflib`](https://github.com/pmezard/go-difflib),
a Go port of Python's `difflib`. `difflib.NewMatcher(a, b).GetOpCodes()` returns
an LCS-based edit script of opcodes; `'i'` (insert) and the target side of `'r'`
(replace) opcodes contribute ADDED lines. Diffing the EMPTY list `A=[]` into
`B=[x,y]` yields a single insert opcode covering every line of `B`, so all lines
are added and the printed added count is `2`.

## Run

    go run .
