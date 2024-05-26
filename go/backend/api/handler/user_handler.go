/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_handler.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:37:43 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/29 04:34:33 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
	"github.com/shiori-42/textbook_change_app/go/backend/service"
)

type UserHandler interface {
	signUp(c echo.Context) error
	logIn(c echo.Context) error
	logOut(c echo.Context) error
}

type userHandler struct{}

func RegisterUserRoutes(e *echo.Echo) {
	h := &userHandler{}
	e.POST("/signup", h.signUp)
	e.POST("/login", h.logIn)
	e.POST("/logout", h.logOut)
}

func (h *userHandler) signUp(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := service.SignUpUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (h *userHandler) logIn(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	storedUser, err := service.LoginUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (h *userHandler) logOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

