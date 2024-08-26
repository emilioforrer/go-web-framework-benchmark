package main

import (
	"fmt"
	"local/go-benchmarks/internal/data"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Create a logrus logger with error level
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.ErrorLevel)

	// Create a new gin router without the default logger middleware
	router := gin.New()

	// Define the route handler
	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", data.Get())
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:8000")
	err := router.Run(":8000")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
