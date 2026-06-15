# 0446 — Indexes

Uses the pure-Go `modernc.org/sqlite` driver through the standard
`database/sql` package on an in-memory database. It creates a `products`
table, inserts three rows, then executes a real `CREATE INDEX idx_sku ON
products(sku)` statement and performs an indexed lookup with
`select price from products where sku=?` bound to `'B'`, printing the
matching price. The connection pool is pinned to one connection
(`SetMaxOpenConns(1)`) because an in-memory database is per-connection.

## Run

    go run .
