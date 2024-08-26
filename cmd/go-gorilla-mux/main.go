package main

import (
	"fmt"
	"local/go-benchmarks/internal/data"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define the handler function and route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data.Get())
	})

	// Start the server using the router
	fmt.Println("Server is running on http://localhost:8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
