package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	if _, err := db.Exec("create table users(id integer, name text)"); err != nil {
		log.Fatal(err)
	}

	users := []struct {
		id   int
		name string
	}{{1, "alice"}, {2, "bob"}, {3, "carol"}}
	for _, u := range users {
		if _, err := db.Exec("insert into users values(?,?)", u.id, u.name); err != nil {
			log.Fatal(err)
		}
	}

	rows, err := db.Query("select name from users order by id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
