# 0008 — Strings

Given `name = "world"`, print a greeting, the name in uppercase, and its
length. Case conversion lives in the `strings` package (`strings.ToUpper`), not
on the string itself. `len(s)` returns the number of **bytes**, not characters —
for the rune (code-point) count use `utf8.RuneCountInString`. For ASCII the two
are equal.

## Run

    go run .
