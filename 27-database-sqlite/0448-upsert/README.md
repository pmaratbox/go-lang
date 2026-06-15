# 0448 — Upsert

Demonstrates an upsert (insert-or-update) against an in-memory SQLite database
using the pure-Go `modernc.org/sqlite` driver through the standard
`database/sql` package. After inserting `apple` with quantity 5, it runs an
`insert ... on conflict(item) do update set qty=qty+excluded.qty` statement so a
second `apple` adds to the existing quantity (10) while a new `banana` is simply
inserted. The connection pool is pinned to one connection
(`SetMaxOpenConns(1)`) because an in-memory database is per-connection.

## Run

    go run .
