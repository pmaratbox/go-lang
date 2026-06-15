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

	if _, err := db.Exec("update users set name='robert' where id=2"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("delete from users where id=1"); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id,name from users order by id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
