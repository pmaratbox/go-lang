package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", ":memory:")
	db.SetMaxOpenConns(1)

	db.Exec("create table t(n integer)")

	// First transaction: insert 1 and 2, then commit.
	tx1, _ := db.Begin()
	tx1.Exec("insert into t values(?)", 1)
	tx1.Exec("insert into t values(?)", 2)
	tx1.Commit()

	// Second transaction: insert 3, then roll back.
	tx2, _ := db.Begin()
	tx2.Exec("insert into t values(?)", 3)
	tx2.Rollback()

	rows, _ := db.Query("select n from t order by n")
	for rows.Next() {
		var n int
		rows.Scan(&n)
		fmt.Println(n)
	}
}
