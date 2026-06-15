# 0452 — Limit

Defines a `User` model (id, name, age) and uses the [GORM](https://gorm.io)
ORM over an in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite`
driver). `db.AutoMigrate` creates the schema from the model and `db.Create`
inserts the rows. The query chains GORM's `Order("age DESC")` and `Limit(2)`
builder methods before `Find` to take only the top two rows by descending age —
all through GORM's API rather than raw SQL. Each user's name is printed.

## Run

    go run .
