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

func NewHandlerBook(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) Index(c *gin.Context) {
	books, err := h.bookService.Index()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, item := range books {
		booksResponse = append(booksResponse, bookResponse(item))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) Store(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Store(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func (h *bookHandler) Show(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	book, err := h.bookService.Show(id)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse(book),
	})
}

func bookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
