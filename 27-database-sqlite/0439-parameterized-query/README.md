# 0439 — Parameterized query

This lesson creates an in-memory SQLite `users` table, inserts three rows, then runs `select name from users where id=?` with the value `2` supplied as a bound query parameter instead of string interpolation. It uses Go's `database/sql` package with the pure-Go `modernc.org/sqlite` driver, binding the parameter through `QueryRow`.

## Run

    go run .
