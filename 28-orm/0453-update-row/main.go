package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})

	db.Create(&[]User{
		{ID: 1, Name: "alice", Age: 30},
		{ID: 2, Name: "bob", Age: 25},
		{ID: 3, Name: "carol", Age: 35},
	})

	// Load bob, change his age, and save the row back through GORM.
	var bob User
	db.Where("name = ?", "bob").First(&bob)
	bob.Age = 40
	db.Save(&bob)

	var users []User
	db.Where("age >= ?", 35).Order("id").Find(&users)
	for _, u := range users {
		fmt.Println(u.Name, u.Age)
	}
}
