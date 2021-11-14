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

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	v1 := router.Group("/v1")

	// routing untuk createBook
	v1.POST("/books", bookHandler.CreateBook)
	// buat routing untuk getBooks
	v1.GET("/books", bookHandler.GetBooks)
	// buat routing untuk getBook berdasarkan Id
	v1.GET("/books/:id", bookHandler.GetBook)
	// routing untuk updateBook
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	// routing untuk deleteBook
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
