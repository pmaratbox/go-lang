# 0445 — Transactions

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package and creates a table `t`.
It runs two transactions via the driver's real transaction control: the
first inserts `1` and `2` and calls `Commit`, the second inserts `3` and
calls `Rollback`. A final `select n from t order by n` prints each surviving
value, showing that the rolled-back `3` never persisted. The connection pool
is pinned to one connection (`SetMaxOpenConns(1)`) because an in-memory
database is per-connection.

## Run

    go run .
