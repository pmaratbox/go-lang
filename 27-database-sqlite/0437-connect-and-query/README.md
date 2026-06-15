# 0437 — Connect & query

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package, then runs the single
query `select 42` and prints the scalar integer result. The connection pool
is pinned to one connection (`SetMaxOpenConns(1)`) because an in-memory
database is per-connection.

## Run

    go run .
