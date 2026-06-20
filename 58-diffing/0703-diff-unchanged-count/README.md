# 0703 — Count unchanged lines

Uses the `github.com/pmezard/go-difflib/difflib` library to diff two fixed
line-lists, `A=[apple,banana,cherry]` against `B=[apple,blueberry,cherry,date]`.
The `SequenceMatcher` opcodes classify each block; the `equal` blocks are the
lines that appear unchanged in both sequences (`apple`, `cherry`). The program
sums the length of those blocks and prints the unchanged count, which is `2`.

## Run

    go run .
