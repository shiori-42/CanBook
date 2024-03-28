/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   item_handler.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/21 11:59:11 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/29 05:16:07 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	// echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/repository"
	"github.com/shiori-42/textbook_change_app/go/backend/service"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RegisterItemRoutes(e *echo.Echo) {
	i := e.Group("/items")
	// i.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:  []byte(os.Getenv("SECRET")),
	// 	TokenLookup: "header:Authorization",
	// }))
	i.Use(AuthMiddleware)
	i.GET("", getMyItems) //my page
	i.GET("/:itemId", getItemByID)
	i.POST("", createItem)
	i.PUT("/:itemId", updateItem)
	i.DELETE("/:itemId", deleteItem)
	e.GET("/search", searchItems)
	e.GET("/searchcollege", searchItemsByCollege)
	e.GET("alluseritems", getAllUserItems) //for no login user
	e.GET("/images/:imageFilename", getImg)
}

func getMyItems(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	items, err := repository.GetMyItems(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func getAllUserItems(c echo.Context) error {
	items, err := repository.GetAllUserItems()
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

func getImg(c echo.Context) error {
	storedDir := "../images/"
	imgPath := path.Join(storedDir, c.Param("imageFilename"))
	if _, err := os.Stat(imgPath); err != nil {
		c.Logger().Errorf("Image not found: %s%s", imgPath, imgPath)
		imgPath = path.Join(storedDir, "default.jpg")
	}
	return c.File(imgPath)
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

func searchItemsByCollege(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	items, err := service.SearchItemsByCollege(keyword)
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
