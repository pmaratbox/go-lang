package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// User is the GORM model that maps to the users table.
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})

	db.Create(&[]User{
		{1, "alice", 30},
		{2, "bob", 25},
		{3, "carol", 35},
	})

	// Sort rows by age ascending using the query builder's Order clause.
	var users []User
	db.Order("age asc").Find(&users)

	for _, u := range users {
		fmt.Println(u.Name)
	}
}
