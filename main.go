package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=pustakaapi password=pustakaapi dbname=pustakaapi port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(book.Book{})

	// book := book.Book{}
	// book.Title = "Man Tiger"
	// book.Price = 9000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "Ini adalah yang sangat bagus dari Eka Kurniawan"

	// CREATE
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// GET
	// buat variable book dulu untuk menampung data book yang diambil dari database
	var book book.Book

	// lalu kita isi variable book atau memparsing data dari yang diambil dari db
	// First artinya yang diambil adalah data book yang ada pada urutan pertama
	// dan parameternya kita ikutnya berupa Primary Key di book nya
	// Debug artinya akan ditampilkan dalam terminal
	err = db.Debug().First(&book, 1).Error
	if err != nil {
		fmt.Println("Error finding book record")
	}
	// jika tidak ada errornya, maka kita tampilkan data book nya
	fmt.Println("Title :", book.Title)
	fmt.Printf("book object %v", book)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
