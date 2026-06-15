# 0442 — Group by

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package, creates a `sales` table,
inserts five rows across two categories, and runs a single
`select category,sum(amount) from sales group by category order by category`
query. Each grouped row is scanned and printed as `category sum`
(space-separated), one per line. The connection pool is pinned to one
connection (`SetMaxOpenConns(1)`) because an in-memory database is
per-connection.

## Run

    go run .
