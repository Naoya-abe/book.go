package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	gorm.Model
	Title string
	Price int
}

// DB migration
func dbInit() {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbInit())")
	}
	db.AutoMigrate(&Book{})
}

// DB Create
func dbInsert(title string, price int) {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbInsert())")
	}
	db.Create(&Book{Title: title, Price: price})
}

// DB Update
func dbUpdate(id int, title string, price int) {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
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
	db, err := gorm.Open("sqlite3", "book.sqlite3")
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
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	var books []Book
	db.Order("created_at desc").Find(&books)
	return books
}

// DB One Get
func dbGetOne(id int) Book {
	db, err := gorm.Open("sqlite3", "book.sqlite3")
	defer db.Close()
	if err != nil {
		panic("You can't open DB (dbGetOne())")
	}
	var book Book
	db.First(&book, id)
	db.Close()
	return book
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	dbInit()

	//Index
	router.GET("/", func(c *gin.Context) {
		books := dbGetAll()
		c.HTML(200, "index.html", gin.H{"books": books})
	})

	//Create
	router.POST("/new", func(c *gin.Context) {
		title := c.PostForm("title")
		p := c.PostForm("price")
		price, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		dbInsert(title, price)
		c.Redirect(302, "/")
	})

	//Edit
	router.GET("/edit/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		book := dbGetOne(id)
		c.HTML(200, "edit.html", gin.H{"book": book})
	})

	//Update
	router.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		title := c.PostForm("title")
		p := c.PostForm("price")
		price, errPrice := strconv.Atoi(p)
		if errPrice != nil {
			panic(errPrice)
		}
		dbUpdate(id, title, price)
		c.Redirect(302, "/")
	})

	//delete
	router.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		dbDelete(id)
		c.Redirect(302, "/")
	})

	//delete_confirm
	router.GET("/delete_confirm/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		book := dbGetOne(id)
		c.HTML(200, "delete.html", gin.H{"book": book})
	})

	router.Run()
}
