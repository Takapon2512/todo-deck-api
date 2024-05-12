package routes

import (
	"todo-deck-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)

	return e
}