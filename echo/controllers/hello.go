package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}
func (controller *HelloController) GetMessage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hallow World!")
}
