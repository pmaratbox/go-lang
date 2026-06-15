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

	var answer int
	if err := db.QueryRow("select 42").Scan(&answer); err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
