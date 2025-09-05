package main

import (
	"fmt"
	"log"

	"book.com/internal/db"
	"book.com/internal/handlers"
	"book.com/internal/models"
	"book.com/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // pre-existing logger middleware

	db.Init()
	if db.DB == nil {
		log.Fatal("Db not initialized")
	}
	fmt.Println("App started successfully after DB connection.")

	err := db.DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database schema")
	}
	//go-routine to save the data into db that is present in channel.
	go handlers.SaveToDB()
	// Routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}
