# 0454 — Delete a row

Uses the [GORM](https://gorm.io) ORM over an in-memory SQLite database
(via the pure-Go `github.com/glebarez/sqlite` driver). The schema is
created from the `User` model with `AutoMigrate`, rows are inserted with
`db.Create`, and the row with `id=1` is removed with `db.Delete(&User{}, 1)`.
The remaining users are read back with `db.Order("id").Find(&users)` so the
output is deterministic — all through GORM's API rather than raw SQL.

## Run

    go run .
