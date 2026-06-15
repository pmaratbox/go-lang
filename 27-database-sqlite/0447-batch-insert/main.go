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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into t values(?)")
	for i := 1; i <= 1000; i++ {
		stmt.Exec(i)
	}
	stmt.Close()
	tx.Commit()

	var count int
	db.QueryRow("select count(*) from t").Scan(&count)
	fmt.Println(count)
}
