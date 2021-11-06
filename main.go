package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Bayu Bagus Bagaswara",
			"bio":  "A Software engineer & content creator",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "Belajar Golang",
		})
	})

	// kita bisa mengubah port untuk localhost
	router.Run(":8081")

}
