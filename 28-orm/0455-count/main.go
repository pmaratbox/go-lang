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
	db.Create(&[]User{{1, "alice", 30}, {2, "bob", 25}, {3, "carol", 35}})

	var count int64
	db.Model(&User{}).Count(&count)
	fmt.Println(count)
}
