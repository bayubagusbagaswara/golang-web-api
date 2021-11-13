package main

import (
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

	// bookRepository := book.NewRepository(db)
	bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookFileRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}
