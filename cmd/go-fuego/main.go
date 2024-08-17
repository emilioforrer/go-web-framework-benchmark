package main

import "github.com/go-fuego/fuego"

func main() {
	s := fuego.NewServer(fuego.WithAddr(":8000"))

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	s.Run()
}
