# 0441 — Aggregate functions

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package, creates a one-column
table, inserts five integer rows, and runs a single
`select count(*),sum(amount),min(amount),max(amount)` query. The four
aggregate results are scanned from the row and printed each on its own line.
The connection pool is pinned to one connection (`SetMaxOpenConns(1)`)
because an in-memory database is per-connection.

## Run

    go run .
