# 0540 — Many (repetition)

Using the `github.com/vektah/goparsify` parser-combinator library, this lesson
builds a parser with the `Many` combinator, which applies its inner parser
zero-or-more times. Here it repeats the char `'a'` parser against the fixed
input `"aaaa"`; each match becomes a child of the result, and `Map` reports how
many were parsed (`4`) via `goparsify.Run`.

## Run

    go run .
