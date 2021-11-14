package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// kita buat endpoint baru untuk mengambil data book yang jumlahnya lebih dari satu
func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// kita ubah menjadi BooksResponse
	var booksResponse []book.BookResponse

	// lakukan looping
	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}
		booksResponse = append(booksResponse, bookResponse)
	}

	// jika ada data book nya, maka kita kembalikan data book nya
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// buat handler untuk GetBook (mengambil single data book berdasarkan ID)
func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// konversi dulu menjadi BookResponse
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	// jika data book dengan id yang dicari ada di database, maka balikkan data booknya berupa bookResponse
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// buat handler untuk Create Book
func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	b, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// konversi dulu menjadi BookResponse
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// dapatkan id nya dahulu
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// konversi dulu menjadi BookResponse
	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}
