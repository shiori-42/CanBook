package handler

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/service"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

func RegisterItemRoutes(e *echo.Echo) {
	i := e.Group("/items")
	i.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	i.Use(AuthMiddleware)
	i.GET("", getAllItems)
	i.GET("/:itemId", getItemByID)
	i.POST("", createItem)
	i.PUT("/:itemId", updateItem)
	i.DELETE("/:itemId", deleteItem)
	e.GET("/search", searchItems)
}

func getAllItems(c echo.Context) error {
	items, err := repository.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func getItemByID(c echo.Context) error {
    itemID := c.Param("itemId")
    item, err := service.GetItemByID(itemID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
    }
    return c.JSON(http.StatusOK, item)
}

func createItem(c echo.Context) error {
	item, err := service.CreateItem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userID := c.Get("user_id").(uint)
	item.UserID = userID

	if err := repository.CreateItem(&item); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, item)
}

func updateItem(c echo.Context) error {
	item, err := service.UpdateItem(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userID := c.Get("user_id").(uint)
	item.UserID = userID

	return c.JSON(http.StatusOK, item)
}

func deleteItem(c echo.Context) error {
	itemID := c.Param("itemId")
	userID := c.Get("user_id").(uint)

	if err := service.DeleteItem(itemID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func searchItems(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	items, err := service.SearchItems(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}