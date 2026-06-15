# 0458 — Join

Defines a `User` model (id, name) with a `has many` `Posts` association and a
`Post` model (id, user_id, title), using the [GORM](https://gorm.io) ORM over an
in-memory SQLite database (the pure-Go `github.com/glebarez/sqlite` driver).
`db.AutoMigrate` creates both tables from the models and `db.Create` inserts the
rows. The posts are joined to their users with GORM's query API —
`db.Model(&Post{}).Select(...).Joins("JOIN users ...").Order("users.name, posts.title").Scan(&rows)`
— and each `name title` pair is printed in name, title order (alice hello,
alice world, bob hi).

## Run

    go run .
