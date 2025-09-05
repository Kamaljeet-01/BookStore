package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// setup global logger
var logger *log.Logger

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file: ", err)
	}
	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("token") == "authOK") {
			c.AbortWithStatusJSON(500, gin.H{
				"message ": "Token is not valid or not present",
			})
			return
		}
		c.Next()
	}
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Key", "Value")
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()

		logger.Printf("[%d] %s %s %s (%s)",
			status,
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			duration,
		)
	}
}
