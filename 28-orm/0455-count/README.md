# 0455 — Count

Defines a `User` model and migrates it into an in-memory SQLite database with
the GORM ORM (`gorm.io/gorm`) over the pure-Go `github.com/glebarez/sqlite`
driver. After inserting three rows with `db.Create`, it counts all users via
the query API's `db.Model(&User{}).Count(&count)` aggregate and prints the
result (`3`).

## Run

    go run .
