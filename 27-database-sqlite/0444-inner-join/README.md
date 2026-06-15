# 0444 — Inner join

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package, creates `users` and
`orders` tables, and inserts a few rows. It then runs an inner join
(`select u.name,o.item from orders o join users u on u.id=o.user_id order by
u.name,o.item`) and prints each matched pair as `name item`. The connection
pool is pinned to one connection (`SetMaxOpenConns(1)`) because an in-memory
database is per-connection.

## Run

    go run .
