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

	var total int64
	db.Model(&User{}).Select("sum(age)").Scan(&total)
	fmt.Println(total)
}
