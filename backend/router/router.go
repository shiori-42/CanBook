package router

import (
	"os"
	// "net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"github.com/shiori-42/textbook_change_app/controller"
)

func NewRouter(uc controller.IUserController, lc controller.IListingController) *echo.Echo {
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
	// 		echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
	// 	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	// 	AllowCredentials: true,
	// }))
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	CookiePath:     "/",
	// 	CookieDomain:   os.Getenv("API_DOMAIN"),
	// 	CookieHTTPOnly: true,
	// 	CookieSameSite: http.SameSiteNoneMode,
	// 	//CookieSameSite: http.SameSiteDefaultMode,
	// 	//CookieMaxAge:   60,
	// }))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	// e.GET("/csrf", uc.CsrfToken)
	l := e.Group("/listing")
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.GET("", lc.GetAllMyListings)
	l.GET("/:listingId", lc.GetMyListingById)
	l.POST("", lc.CreateListing)
	l.PUT("/:listingId", lc.UpdateListing)
	l.DELETE("/:listingId", lc.DeleteListing)
	return e
}
