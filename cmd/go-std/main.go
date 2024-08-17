package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
