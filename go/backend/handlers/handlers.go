package main

import (
	"io"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shiori-42/textbook_change_app/go/backend/api/handler"
	// "github.com/shiori-42/textbook_change_app/go/backend/handler"
)

type JetRenderer struct {
	views *jet.Set
}

func NewJetRenderer() *JetRenderer {
	return &JetRenderer{
		views: jet.NewSet(
			jet.NewOSFileSystemLoader("./html"),
			jet.InDevelopmentMode(),
		),
	}
}

func (r *JetRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	view, err := r.views.GetTemplate(name)
	if err != nil {
		return err
	}

	vars, ok := data.(jet.VarMap)
	if !ok {
		vars = make(jet.VarMap)
	}

	return view.Execute(w, vars, nil)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = NewJetRenderer()
	// Register WebSocket routes
	handler.RegisterWebSocketRoutes(e)

	e.GET("/", func(c echo.Context) error {
		data := make(jet.VarMap)
		data.Set("title", "Home Page")
		return c.Render(http.StatusOK, "home.jet", data)
	})
	e.Logger.Fatal(e.Start(":8080"))

}
