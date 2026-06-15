package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Posts []Post
}

type Post struct {
	ID     int
	UserID int
	Title  string
}

func main() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Post{})

	db.Create(&[]User{
		{ID: 1, Name: "alice"},
		{ID: 2, Name: "bob"},
	})
	db.Create(&[]Post{
		{ID: 1, UserID: 1, Title: "hello"},
		{ID: 2, UserID: 1, Title: "world"},
		{ID: 3, UserID: 2, Title: "hi"},
	})

	// Load each user together with its Posts association (ordered by title),
	// using GORM's relationship preloading rather than a hand-written join.
	var users []User
	db.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("title")
	}).Order("name").Find(&users)

	for _, u := range users {
		for _, p := range u.Posts {
			fmt.Printf("%s %s\n", u.Name, p.Title)
		}
	}
}
