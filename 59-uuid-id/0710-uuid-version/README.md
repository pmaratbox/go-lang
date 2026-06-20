# 0710 — UUID version

Read a UUID's version number with the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid). The version is
encoded in the UUID itself, so once `uuid.MustParse` parses
`550e8400-e29b-41d4-a716-446655440000` we call `u.Version()` to recover it.
Note that `Version`'s `String()` is `"VERSION_4"`, so we wrap it in `int()` to
print just the bare number `4`.

## Run

    go run .
