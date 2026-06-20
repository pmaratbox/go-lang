# 0700 — Count removed lines

Uses the `github.com/pmezard/go-difflib/difflib` library to diff two fixed
line-lists, `A=[apple,banana,cherry]` against `B=[apple,blueberry,cherry,date]`.
The `SequenceMatcher` opcodes classify each block; `delete` and the source side
of `replace` deltas are the removed lines. The program prints the number of
removed lines (`banana`), which is `1`.

## Run

    go run .
