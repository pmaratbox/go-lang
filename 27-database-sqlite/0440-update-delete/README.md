# 0440 — Update & delete

Seeds an in-memory SQLite database (three users) through the pure-Go
`modernc.org/sqlite` driver and the standard `database/sql` package, then
mutates it: an `UPDATE` renames the row with `id=2` to `robert` and a
`DELETE` removes the row with `id=1`. It re-reads the table with
`select id,name from users order by id` and prints each surviving row as
`id name`. The pool is pinned to one connection (`SetMaxOpenConns(1)`)
because an in-memory database is per-connection.

## Run

    go run .
