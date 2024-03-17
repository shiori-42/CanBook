package router

import (
	"os"

	"github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/controller"
)

func NewRouter(uc controller.IUserController, lc controller.IListingController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	l := e.Group("/listing")
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.GET("", lc.GetAllMyListings)
	l.GET("/:id", lc.GetMyListingById)
	l.POST("", lc.CreateListing)
	l.PUT("/:id", lc.UpdateListing)
	l.DELETE("/:id", lc.DeleteListing)
	return e
}
