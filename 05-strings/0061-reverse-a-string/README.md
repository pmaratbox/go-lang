# 0061 — Reverse a String

Reverse the string `abc` character by character and print the result: `cba`. A Go `string` is UTF-8 bytes, so it is converted to `[]rune` first; reversing the runes with a two-pointer swap keeps multi-byte characters intact (reversing bytes would corrupt them).

## Run

    go run .
