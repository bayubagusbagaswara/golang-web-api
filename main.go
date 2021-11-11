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

	// kita akan melakukan operasi CRUD untuk memanipulasi data
	// kita bikin object dari struct book
	book := book.Book{}
	// isi property pada struct book
	book.Title = "Man Tiger"
	book.Price = 9000
	book.Discount = 10
	book.Rating = 5
	book.Description = "Ini adalah yang sangat bagus dari Eka Kurniawan"

	// lalu kita simpan data struct diatas, dan balikannya adalah error, yang nantinya dicek apakah ada error atau tidak
	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("Error creating book record")
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
