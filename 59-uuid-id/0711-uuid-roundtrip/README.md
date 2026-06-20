# 0711 — Parse and format

Round-trip a UUID through the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid). We parse the
uppercase string `550E8400-E29B-41D4-A716-446655440000` with `uuid.MustParse`
(which accepts mixed case) and print it back with `String()`, which always
renders the canonical lowercase form `550e8400-e29b-41d4-a716-446655440000`.

## Run

    go run .
