# 0449 — Define model & insert

Defines a `User` model (id, name, age) and uses the [GORM](https://gorm.io)
ORM over an in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite`
driver). `db.AutoMigrate` creates the schema from the model, `db.Create`
inserts the rows, and `db.Order("id").Find` reads them back ordered by id —
all through GORM's API rather than raw SQL. Each user's name is printed.

## Run

    go run .
