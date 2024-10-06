package main

import (
	"local/go-benchmarks/internal/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data.Get())
	})
	http.ListenAndServe(":8000", r)
}
