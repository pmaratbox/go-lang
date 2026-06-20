# 0705 — Identical inputs

Diff two fixed line-lists with the real diff library
[`github.com/pmezard/go-difflib/difflib`](https://github.com/pmezard/go-difflib),
a Go port of Python's `difflib`. `difflib.NewMatcher(a, b).GetOpCodes()` returns
an LCS-based edit script of opcodes; insert (`'i'`) and the target side of
replace (`'r'`) opcodes add lines, while delete (`'d'`) and the source side of
replace remove lines. Diffing `A=[apple,banana,cherry]` against itself yields a
single `'e'` (equal) opcode, so the printed `<added> <removed>` counts are `0 0`.

## Run

    go run .
