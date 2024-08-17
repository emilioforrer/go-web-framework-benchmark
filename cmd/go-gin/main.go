package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	router := gin.Default()

	// Define the route handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:8000")
	err := router.Run(":8000")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
