# 0453 — Update a row

Modifies a persisted entity using the [GORM](https://gorm.io) ORM over an
in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite` driver).
After inserting three users, `db.Where(...).First` loads bob, his `Age` field
is changed, and `db.Save` writes the whole row back. A follow-up
`db.Where("age >= ?", 35).Order("id").Find` reads the matching rows — all
through GORM's API rather than raw SQL — printing each `name age`.

## Run

    go run .
