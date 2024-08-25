// package main

// import (
// 	"net/http"
// 	"log"

// 	"github.com/CloudyKit/jet/v6"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"github.com/shiori-42/textbook_change_app/go/backend/api/handler"
// 	// "github.com/shiori-42/textbook_change_app/go/backend/handler"
// )

// type JetRenderer struct {
// 	views *jet.Set
// }

// func NewJetRenderer() *JetRenderer {
// 	return &JetRenderer{
// 		views: jet.NewSet(
// 			jet.NewOSFileSystemLoader("./html"),
// 			jet.InDevelopmentMode(),
// 		),
// 	}
// }

// func Home(c echo.Context)error{
// 	err:=renderPage(c,"home,jet",nil)
// 	if err!=nil{
// 		log.Println(err)
// 	}
// 	return err
// }

// func renderPage(c echo.Context, tmpl string, data jet.VarMap) error {
// 	view, err := r.views.GetTemplate(tmpl)
// 	if err != nil {
// 		return err
// 	}

// 	err=view.Execute(c.Response().Writer,data,nil)
// 	if err!=nil{
// 		return err
// 	}
// 	return nil
// }

// func main() {
// 	e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.Renderer = NewJetRenderer()
// 	// Register WebSocket routes
// 	handler.RegisterWebSocketRoutes(e)

// 	e.GET("/", func(c echo.Context) error {
// 		data := make(jet.VarMap)
// 		data.Set("title", "Home Page")
// 		return c.Render(http.StatusOK, "home.jet", data)
// 	})
// 	e.Logger.Fatal(e.Start(":8080"))

// }
