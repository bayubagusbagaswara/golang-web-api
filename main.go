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

	// kita tidak akan langsung berhubungan dengan database
	// tapi kita akan buat Repository Layer untuk berkomunikasi dengan database

	db.AutoMigrate(book.Book{})

	// kita parsing object db nya
	bookRepository := book.NewRepository(db)

	// kita panggil function yang ada di repository
	// books, err := bookRepository.FindAll()
	// kita buktikan bahwa books ada isi datanya
	// for _, book := range books {
	// 	fmt.Println("Title: ", book.Title)
	// }

	// function findById
	// book, err := bookRepository.FindById(1)
	// fmt.Println("Title :", book.Title)

	// function Create
	// buat dulu data untuk struct book
	book := book.Book{
		Title:       "Laskar Pelangi",
		Description: "Buku anak petualangan",
		Price:       95000,
		Rating:      4,
		Discount:    0,
	}
	// lalu save
	bookRepository.Create(book)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
