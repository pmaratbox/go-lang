# 0712 — UUIDv5 of another name

Generate a name-based UUIDv5 with the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid) using
`uuid.NewSHA1(uuid.NameSpaceDNS, []byte("test.com"))`. UUIDv5 hashes the
(namespace, name) pair with SHA-1, so it is fully deterministic — yet the result
depends on the name: hashing `test.com` under the DNS namespace
(`6ba7b810-9dad-11d1-80b4-00c04fd430c8`) produces a different UUID than
`example.com` did. The value is never hardcoded; it is computed each run.

## Run

    go run .
