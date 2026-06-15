package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID       int
	Category string
	Price    int
}

func main() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Product{})

	db.Create(&[]Product{
		{ID: 1, Category: "a", Price: 10},
		{ID: 2, Category: "b", Price: 20},
		{ID: 3, Category: "a", Price: 30},
	})

	type Row struct {
		Category string
		Total    int
	}
	var rows []Row
	db.Model(&Product{}).
		Select("category, sum(price) as total").
		Group("category").
		Order("category").
		Scan(&rows)

	for _, r := range rows {
		fmt.Printf("%s %d\n", r.Category, r.Total)
	}
}
