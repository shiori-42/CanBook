package router

import (
	"github.com/shiori-42/textbook_change_app/controller"

	"github.com/labstack/echo/v4"
)

func New(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	return e
}
