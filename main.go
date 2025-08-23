package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getFromRequest(book *Book, c *gin.Context) bool {
	if err := c.BindJSON(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return false
	}
	return true
}

func addBook(c *gin.Context) {
	var book Book
	if !getFromRequest(&book, c) {
		return
	}

	Shelf = append(Shelf, book)
	c.JSON(201, gin.H{
		"message": "book store sucessfully",
	})
}
func seeBooks(c *gin.Context) {
	c.JSON(http.StatusAccepted, Shelf) //response
}

func update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var book UpdatedBook
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	//find book index
	idx := Find(id)
	if idx != -1 {
		Shelf[idx].Name = book.Name
		Shelf[idx].Id = id
		Shelf[idx].Price = book.Price      //updated the Book Details
		c.JSON(http.StatusAccepted, Shelf) //response after updation
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Book you want to update is not in the Store",
		})
	}
}
func delete(c *gin.Context) {
	var book Book
	if !getFromRequest(&book, c) {
		return
	}
	idx := Find(book.Id)
	if idx != -1 {
		Shelf = append(Shelf[:idx], Shelf[idx+1:]...)
		c.JSON(http.StatusAccepted, gin.H{"Book is deleted from the store": Shelf})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message ": "This book is not found in the store,"})
	}

}
func setRoutes(r *gin.Engine) {

	book := r.Group("/books")
	{
		book.GET("/see", seeBooks)
		book.POST("/add", addBook)
		book.PUT("/update/:id", update)
		book.DELETE("/delete", delete)
	}
}
func main() {
	r := gin.Default()
	setRoutes(r)
	r.Run(":8080")
}
