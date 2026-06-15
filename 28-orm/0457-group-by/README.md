# 0457 — Group by

Defines a `Product` model (id, category, price) and uses the [GORM](https://gorm.io)
ORM over an in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite`
driver). `db.AutoMigrate` creates the schema from the model and `db.Create`
inserts the rows. The per-category price total is computed with GORM's
query-builder API — `db.Model(&Product{}).Select("category, sum(price) as total").Group("category").Order("category").Scan(&rows)`
— rather than a raw SQL string, and each `category sum` pair is printed in
category order (a 40, b 20).

## Run

    go run .
