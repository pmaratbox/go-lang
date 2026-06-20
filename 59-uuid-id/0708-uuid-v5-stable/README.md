# 0708 — UUIDv5 is stable

Generate a name-based UUIDv5 twice with the real UUID library
[`github.com/google/uuid`](https://github.com/google/uuid) using
`uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))`. UUIDv5 hashes the
(namespace, name) pair with SHA-1, so it is fully deterministic — unlike the
random UUIDv4. Computing it twice from the DNS namespace
(`6ba7b810-9dad-11d1-80b4-00c04fd430c8`) and the same name yields byte-identical
values, so comparing the two results prints `true`.

## Run

    go run .
