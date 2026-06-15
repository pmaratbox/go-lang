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

	if _, err := db.Exec("create table t(amount integer)"); err != nil {
		log.Fatal(err)
	}
	for _, amount := range []int{10, 20, 30, 40, 50} {
		if _, err := db.Exec("insert into t values(?)", amount); err != nil {
			log.Fatal(err)
		}
	}

	var count, sum, min, max int
	row := db.QueryRow("select count(*),sum(amount),min(amount),max(amount) from t")
	if err := row.Scan(&count, &sum, &min, &max); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	fmt.Println(sum)
	fmt.Println(min)
	fmt.Println(max)
}
