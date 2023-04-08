package infra

import (
	"app/controllers"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	return &Server{
		Echo: echo.New(),
	}
}

func (server *Server) Start() {
	server.Echo.GET("/", controllers.NewHelloController().GetMessage)

	c := jaegertracing.New(server.Echo, nil)
	defer c.Close()

	server.Echo.Start(":1323")
}
