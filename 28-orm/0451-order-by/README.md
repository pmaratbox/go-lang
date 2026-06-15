# 0451 — Order by

Uses the [GORM](https://gorm.io) ORM over an in-memory SQLite database
(via the pure-Go `github.com/glebarez/sqlite` driver). The schema is
created from the `User` model with `AutoMigrate`, rows are inserted with
`db.Create`, and results are sorted with the query builder's
`db.Order("age asc").Find(&users)` clause so the output is deterministic.

## Run

    go run .
