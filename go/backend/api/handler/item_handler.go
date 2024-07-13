/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_handler.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori <shiori@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/21 11:59:11 by shiori0123        #+#    #+#             */
/*   Updated: 2024/07/13 18:25:17 by shiori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
	"github.com/shiori-42/textbook_change_app/go/backend/service"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ItemHandler interface {
	getMyItems(c echo.Context) error
	getAllUserItems(c echo.Context) error
	getItemByID(c echo.Context) error
	getImg(c echo.Context) error
	createItem(c echo.Context) error
	updateItem(c echo.Context) error
	deleteItem(c echo.Context) error
	searchItems(c echo.Context) error
	searchItemsByCollege(c echo.Context) error
}

type itemHandler struct{}

func RegisterItemRoutes(e *echo.Echo) {
	h := &itemHandler{}
	i := e.Group("/items")
	i.Use(AuthMiddleware)
	i.GET("", h.getMyItems) //my page
	i.GET("/:itemId", h.getItemByID)
	i.POST("", h.createItem)
	i.PUT("/:itemId", h.updateItem)
	i.DELETE("/:itemId", h.deleteItem)
	e.GET("/search", h.searchItems)
	e.GET("/searchcollege", h.searchItemsByCollege)
	e.GET("alluseritems", h.getAllUserItems) //for no login user
	e.GET("/images/:imageFilename", h.getImg)
}

func (h *itemHandler) getMyItems(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	items, err := repository.GetMyItems(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func (h *itemHandler) getAllUserItems(c echo.Context) error {
	items, err := repository.GetAllUserItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func (h *itemHandler) getItemByID(c echo.Context) error {
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

func (h *itemHandler) getImg(c echo.Context) error {
	uploadDir := os.Getenv("UPLOAD_DIR")
	imgPath := path.Join(uploadDir, c.Param("imageFilename"))
	if _, err := os.Stat(imgPath); err != nil {
		c.Logger().Errorf("Image not found: %s%s", imgPath, imgPath)
		imgPath = path.Join(uploadDir, "default.jpg")
	}
	return c.File(imgPath)
}

func (h *itemHandler) createItem(c echo.Context) error {
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

func (h *itemHandler) updateItem(c echo.Context) error {
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

func (h *itemHandler) deleteItem(c echo.Context) error {
	itemID := c.Param("itemId")
	userID := c.Get("user_id").(uint)

	if err := service.DeleteItem(itemID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *itemHandler) searchItems(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	items, err := service.SearchItemsByKeyword(keyword)
	if err != nil {
		c.Logger().Errorf("SearchItemsByKeyword error: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func (h *itemHandler) searchItemsByCollege(c echo.Context) error {
	college := c.QueryParam("college")
	items, err := service.SearchItemsByCollege(college)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Authorization header is required"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Authorization header must be in the format 'Bearer {token}'"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", uint(userID))
				return next(c)
			}
		}
		return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid or expired token"})
	}
}
