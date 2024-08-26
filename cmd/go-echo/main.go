package main

import (
	"local/go-benchmarks/internal/data"
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
	return c.Blob(http.StatusOK, "text/plain", data.Get())
}
