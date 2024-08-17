package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	// Create a new router
	router := bunrouter.New(
	// bunrouter.Use(reqlog.NewMiddleware()),
	)

	// Define a route
	router.GET("/", helloHandler)

	// Create an HTTP server
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	// Start the server
	fmt.Println("Server is running on http://localhost:8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Handler function
func helloHandler(w http.ResponseWriter, req bunrouter.Request) error {
	// Write the plain text response
	_, err := w.Write([]byte("Hello, World!"))
	return err
}
