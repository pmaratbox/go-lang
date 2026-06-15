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
	// The in-memory database lives for the lifetime of a single connection,
	// so pin the pool to exactly one connection.
	db.SetMaxOpenConns(1)
	defer db.Close()

	if _, err := db.Exec("create table scores(value integer)"); err != nil {
		log.Fatal(err)
	}
	for _, value := range []int{50, 90, 70, 30, 100, 20} {
		if _, err := db.Exec("insert into scores values(?)", value); err != nil {
			log.Fatal(err)
		}
	}

	rows, err := db.Query("select value from scores order by value desc limit 3")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var value int
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
