package main

import (
	"net/http"
	"io"

	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		vars = jet.VarMap{}
	}

	return view.Execute(w, vars, nil)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = NewJetRenderer()

	e.GET("/", func(c echo.Context) error {
		data := jet.VarMap{
			"title":  jet.Var("Home Page"),
		}
		return c.Render(http.StatusOK, "home.jet", data)
	})

}
