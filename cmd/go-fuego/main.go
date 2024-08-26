package main

import (
	"local/go-benchmarks/internal/data"
	"net/http"

	"github.com/go-fuego/fuego"
)

func main() {
	s := fuego.NewServer(fuego.WithAddr(":8000"))

	// Standard net/http handler with automatic OpenAPI route declaration
	fuego.GetStd(s, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data.Get())
	})

	s.Run()
}
