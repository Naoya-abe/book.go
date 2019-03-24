package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title string
	Price int
}

// DB migration
func dbInit() {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbInit())")
	}
	db.AutoMigrate(&Book{})
}

// DB Create
func dbInsert(title string, price int) {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbInsert())")
	}
	db.Create(&Book{Title: title, Price: price})
}

// DB Update
func dbUpdate(id int, title string, price int) {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbUpdate())")
	}
	var book Book
	db.First(&book, id)
	book.Title = title
	book.Price = price
	db.Save(&book)
}

// DB Delete
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbDelete())")
	}
	var book Book
	db.First(&book, id)
	db.Delete(&book)
}

// DB All Get
func dbGetAll() []Book {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	var books []Book
	db.Order("created_at desc").Find(&books)
	return books
}

// DB One Get
func dbGetOne(id int, title string, price int) {
	db, err := gorm.Open("sqlite3", "books.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbGetOne())")
	}
	var book Book
	db.First(&book, id)
	db.Close()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

}
