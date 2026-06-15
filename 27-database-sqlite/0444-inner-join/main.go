package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", ":memory:")
	db.SetMaxOpenConns(1)

	db.Exec("create table users(id integer, name text)")
	db.Exec("create table orders(user_id integer, item text)")

	users := []struct {
		id   int
		name string
	}{{1, "alice"}, {2, "bob"}}
	for _, u := range users {
		db.Exec("insert into users values(?,?)", u.id, u.name)
	}

	orders := []struct {
		userID int
		item   string
	}{{1, "book"}, {2, "pen"}, {1, "lamp"}}
	for _, o := range orders {
		db.Exec("insert into orders values(?,?)", o.userID, o.item)
	}

	rows, _ := db.Query("select u.name,o.item from orders o join users u on u.id=o.user_id order by u.name,o.item")
	for rows.Next() {
		var name, item string
		rows.Scan(&name, &item)
		fmt.Println(name, item)
	}
}
