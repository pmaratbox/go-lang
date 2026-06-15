# 0538 — Sequence

Using the `github.com/vektah/goparsify` parser-combinator library, this lesson
builds a parser with the `Seq` combinator: it matches the char `'a'` THEN the
char `'b'`, in order. The two matched tokens (held in each child's `Token`) are
combined into a single string with `Map`, and `goparsify.Run` parses the fixed
input `"ab"`.

## Run

    go run .
