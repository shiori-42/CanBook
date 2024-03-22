/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_handler.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/21 11:59:11 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/22 15:51:51 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
	"net/http"
	"os"
	"strconv"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
	"github.com/shiori-42/textbook_change_app/go/backend/service"
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
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid item id"})
	}
	item, err := service.GetItemByID(itemID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

func createItem(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid user id"})
	}
	item, err := service.CreateItem(c, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, item)
}

func updateItem(c echo.Context) error {
	itemIDStr := c.Param("itemId")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid item id"})
	}
	userID := c.Get("user_id").(uint)
	item, err := service.UpdateItem(c, itemID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
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
	items, err := service.SearchItemsByKeyword(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}
