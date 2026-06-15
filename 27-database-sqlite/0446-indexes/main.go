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

	if _, err := db.Exec("create table products(id integer, sku text, price integer)"); err != nil {
		log.Fatal(err)
	}

	products := []struct {
		id    int
		sku   string
		price int
	}{{1, "A", 100}, {2, "B", 200}, {3, "C", 300}}
	for _, p := range products {
		if _, err := db.Exec("insert into products values(?,?,?)", p.id, p.sku, p.price); err != nil {
			log.Fatal(err)
		}
	}

	// Actually execute the CREATE INDEX statement so lookups can use it.
	if _, err := db.Exec("create index idx_sku on products(sku)"); err != nil {
		log.Fatal(err)
	}

	// Indexed lookup with a real bound parameter.
	var price int
	if err := db.QueryRow("select price from products where sku=?", "B").Scan(&price); err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
