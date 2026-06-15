# 0438 — Create table & insert

Create a `users` table in an in-memory SQLite database, insert three rows with prepared-statement parameter binding, then `select name from users order by id` and print each name on its own line. Uses Go's `database/sql` package with the pure-Go `modernc.org/sqlite` driver (driver name `sqlite`).

## Run

    go run .
