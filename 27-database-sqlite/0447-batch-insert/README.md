# 0447 — Batch insert

Opens an in-memory SQLite database with the pure-Go `modernc.org/sqlite`
driver through the standard `database/sql` package, then inserts 1000 rows
(values 1..1000) efficiently inside a single transaction using a prepared
statement (`db.Begin`, `tx.Prepare`, repeated `stmt.Exec`, `tx.Commit`).
It finally runs `select count(*) from t` and prints the row count. The
connection pool is pinned to one connection (`SetMaxOpenConns(1)`) because
an in-memory database is per-connection.

## Run

    go run .
