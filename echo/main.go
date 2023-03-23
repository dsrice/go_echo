package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	server := echo.New()
	server.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World")
	})

	server.Logger.Fatal(server.Start(":1323"))
}
