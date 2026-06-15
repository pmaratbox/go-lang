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

	if _, err := db.Exec("create table sales(category text, amount integer)"); err != nil {
		log.Fatal(err)
	}
	sales := []struct {
		category string
		amount   int
	}{{"a", 10}, {"b", 20}, {"a", 30}, {"b", 5}, {"a", 2}}
	for _, s := range sales {
		if _, err := db.Exec("insert into sales values(?,?)", s.category, s.amount); err != nil {
			log.Fatal(err)
		}
	}

	rows, err := db.Query("select category,sum(amount) from sales group by category order by category")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		var sum int
		if err := rows.Scan(&category, &sum); err != nil {
			log.Fatal(err)
		}
		fmt.Println(category, sum)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
