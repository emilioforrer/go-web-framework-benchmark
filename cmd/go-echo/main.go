package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define a route
	e.GET("/", hello)

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
