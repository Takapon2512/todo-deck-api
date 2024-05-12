package controller

import (
	"net/http"
	"os"
	"time"
	"todo-deck-api/model"
	"todo-deck-api/usecase"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{ uu }
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}