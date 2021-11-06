package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// buat variable object untuk gin
	router := gin.Default()

	// buat root URL, parameter kedua adalah function
	// balikannya adalah data JSON yang sudah kita buat dibawah
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Bayu Bagus Bagaswara",
			"bio":  "A Software engineer & content creator",
		})
	})

	// lalu kita Run routernya
	router.Run()

}
