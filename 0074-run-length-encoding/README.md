# 0074 — Run-Length Encoding

Run-length encode the string `aaabbc` (each run of a repeated character becomes the character followed by its count), printing `a3b2c1`. Working over `[]rune`, the inner loop counts each run and `Printf` with the `%c%d` verbs emits the character followed by its count.

## Run

    go run .
