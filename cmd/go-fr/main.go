package main

import (
	"local/go-benchmarks/internal/data"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

func main() {
	app := gofr.New()

	app.GET("/", func(c *gofr.Context) (interface{}, error) {

		return response.File{
			Content:     data.Get(),
			ContentType: "text/plain",
		}, nil

	})

	app.Run() // listen and serve on localhost:8000
}
