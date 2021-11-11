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

	// UPDATE
	// misal kita ambil dulu data yang tersimpan di database
	var book book.Book

	err = db.Debug().First(&book, 1).Error
	if err != nil {
		fmt.Println("Error finding book record")
	}

	// misal kita ingin update data Title nya
	book.Title = "Man Tiger (Revised edition)"
	// lalu simpan data book yang baru
	err = db.Save(&book).Error
	// lalu kita cek error juga
	if err != nil {
		fmt.Println("Error updating book record")
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
