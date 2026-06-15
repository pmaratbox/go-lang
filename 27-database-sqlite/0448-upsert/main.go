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

	if _, err := db.Exec("create table inv(item text primary key, qty integer)"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("insert into inv values(?,?)", "apple", 5); err != nil {
		log.Fatal(err)
	}

	const upsert = "insert into inv(item, qty) values(?,?) " +
		"on conflict(item) do update set qty=qty+excluded.qty"
	if _, err := db.Exec(upsert, "apple", 5); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(upsert, "banana", 3); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select item, qty from inv order by item")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item string
		var qty int
		if err := rows.Scan(&item, &qty); err != nil {
			log.Fatal(err)
		}
		fmt.Println(item, qty)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
