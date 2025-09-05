package routes

import (
	"book.com/internal/handlers"
	"book.com/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	book := r.Group("/books")
	book.Use(
		middleware.Authenticate(),
		middleware.Logger(),
		middleware.ResponseMiddleware(),
	)
	{
		book.GET("/see", handlers.SeeBooks)
		book.POST("/add", handlers.AddBook)
		book.PUT("/update", handlers.Update)
		book.DELETE("/delete/:id", handlers.Delete)
	}
}
