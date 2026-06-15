# 0456 — Sum aggregate

Defines a `User` model (id, name, age) and uses the [GORM](https://gorm.io)
ORM over an in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite`
driver). `db.AutoMigrate` creates the schema from the model and `db.Create`
inserts the rows. The total of the `age` column is computed with GORM's query
API — `db.Model(&User{}).Select("sum(age)").Scan(&total)` — rather than a raw
SQL string, and the sum (90) is printed.

## Run

    go run .
