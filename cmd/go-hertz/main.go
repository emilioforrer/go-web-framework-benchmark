package main

import (
	"context"
	"local/go-benchmarks/internal/data"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8000"),
		server.WithDisablePrintRoute(true),
	)

	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.Data(consts.StatusOK, "text/plain", data.Get())
	})

	h.Spin()
}
