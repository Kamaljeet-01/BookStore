package handlers

import (
	"log"
	"net/http"

	//"strconv"

	"book.com/internal/db"
	"book.com/internal/models"
	"github.com/gin-gonic/gin"
)

func SaveToDB() {
	for book := range db.BookCh {
		err := db.DB.Create(&book).Error
		if err != nil {
			log.Printf("Failed to insert data because %v", err)
		} else {
			log.Printf("Book stored successfully.\n%+v", book)
		}
	}
}

// helper func to fetch data form user request.
func getFromRequest(book *models.Book, c *gin.Context) bool {
	if err := c.BindJSON(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return false
	}
	return true
}

// handler func for "/add" endpoint
func AddBook(c *gin.Context) {
	var book models.Book
	if !getFromRequest(&book, c) {
		return
	}
	//sending into channel
	db.BookCh <- book

	c.JSON(201, gin.H{
		"message": "Data stored sucessfully",
	})
}

func SeeBooks(c *gin.Context) {
	c.JSON(http.StatusAccepted, models.Shelf)
}

func Update(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := strconv.Atoi(idStr)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
	// 	return
	// }

	// var book models.UpdatedBook
	// if err := c.BindJSON(book); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
	// 	return
	// }

	// idx := models.Find(id)
	// if idx != -1 {
	// 	models.Shelf[idx].Name = book.Name
	// 	models.Shelf[idx].Id = id
	// 	models.Shelf[idx].Price = book.Price
	// 	c.JSON(http.StatusAccepted, models.Shelf)
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Message": "Book you want to update is not in the Store",
	// 	})
	// }
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Unscoped().Delete(&models.Book{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted sucessfully."})
}
