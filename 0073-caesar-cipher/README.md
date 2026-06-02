# 0073 — Caesar Cipher

Encrypt `abc` with a Caesar cipher shifting each letter forward by `1` (wrapping within the alphabet) and print the result: `bcd`. A `rune` is an integer, so `ch - 'a'` gives `0..25`; the shift wraps with `% 26` before adding `'a'` back.

## Run

    go run .
