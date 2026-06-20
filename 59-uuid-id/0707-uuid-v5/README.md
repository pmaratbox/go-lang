# 0707 — UUIDv5 (name-based)

Generate a deterministic UUIDv5 with the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid). UUIDv5 is
*name-based*: it hashes a namespace plus a name with SHA-1, so the same
`(namespace, name)` pair always yields the same UUID (unlike the random v4).
Here `uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))` uses the
standard DNS namespace `6ba7b810-9dad-11d1-80b4-00c04fd430c8` and the name
`example.com`, producing `cfbff0d1-9375-5685-968c-48ce8b15ae17`.

## Run

    go run .
