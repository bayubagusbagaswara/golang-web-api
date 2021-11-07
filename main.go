package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)

	router.POST("/books", postBooksHandler)

	router.Run(":8081")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Bayu Bagus Bagaswara",
		"bio":  "A Software engineer & content creator",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Golang",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// kita buat sebuah struct untuk menangkap data JSON yang akan dikirimkan dari client
type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	// kita akan mengirim data title dan price dari sebuah book
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	// kalau tidak ada error, kita balikan status OK dan data hasil POSTnya
	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})

}
