package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iswarmondal/REST_Go/models"
)

var Books = []models.Book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}

func createBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, Books)
}

func getBookById(id string) (*models.Book, error) {
	for i, b := range Books {
		if b.Id == id {
			return &Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	theBook, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, theBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query parameter id"})
		return
	}

	theBook, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "the book is not found"})
		return
	}

	if theBook.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book is not available"})
		return
	}

	theBook.Quantity -= 1
	c.IndentedJSON(http.StatusOK, theBook)
}

func checkinBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query parameter id"})
		return
	}

	theBook, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "the book is not found"})
		return
	}

	theBook.Quantity += 1
	c.IndentedJSON(http.StatusOK, theBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.POST("/add-book", createBook)
	router.GET("/book/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/checkin", checkinBook)
	router.Run("localhost:9988")
}
