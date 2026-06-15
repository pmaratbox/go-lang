# 0443 — Order by & limit

Creates an in-memory SQLite `scores` table, inserts six integers, then runs `select value from scores order by value desc limit 3` to sort descending and take the top rows. Uses the real `modernc.org/sqlite` driver via the standard `database/sql` API, iterating the result set with `rows.Next`/`rows.Scan` and printing each value on its own line.

## Run

    go run .
