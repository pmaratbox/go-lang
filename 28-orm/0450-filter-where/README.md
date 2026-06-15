# 0450 — Filter with where

Define a `User` model and let GORM build the schema in an in-memory SQLite database (`github.com/glebarez/sqlite` driver), insert three rows with `db.Create`, then query the matching rows with the query-builder's `db.Where("age >= ?", 30).Order("id").Find(&users)` API and print each name. Uses the `gorm.io/gorm` ORM.

## Run

    go run .
