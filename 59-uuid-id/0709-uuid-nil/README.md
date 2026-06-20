# 0709 — Nil UUID

Print the nil (all-zero) UUID with the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid). The library
exposes the special all-zero value as the constant `uuid.Nil`; it is the
canonical "empty" UUID `00000000-0000-0000-0000-000000000000`, often used as a
sentinel for "no UUID". Here we simply print `uuid.Nil` rather than hardcoding
the zero string.

## Run

    go run .
