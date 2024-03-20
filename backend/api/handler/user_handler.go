/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   user_handler.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 16:37:43 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 21:12:00 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handler

import (
    "net/http"
    "os"
    "time"
    "fmt"

    "github.com/golang-jwt/jwt/v4"
    "github.com/labstack/echo/v4"
    "github.com/shiori-42/textbook_change_app/model"
    "github.com/shiori-42/textbook_change_app/service"
)

func RegisterUserRoutes(e *echo.Echo) {
    e.POST("/signup", signUp)
    e.POST("/login", logIn)
    e.POST("/logout", logOut)
}

func signUp(c echo.Context) error {
    var user model.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    resUser, err := service.SignUpUser(user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, resUser)
}

func logIn(c echo.Context) error {
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

    return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func logOut(c echo.Context) error {
    cookie := new(http.Cookie)
    cookie.Name = "token"
    cookie.Value = ""
    cookie.Expires = time.Now()
    cookie.Path = "/"
    cookie.Domain = os.Getenv("API_DOMAIN")
    cookie.HttpOnly = true
    cookie.SameSite = http.SameSiteNoneMode

    c.SetCookie(cookie)
    return c.NoContent(http.StatusOK)
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString, err := c.Cookie("token")
        if err != nil {
            return c.JSON(http.StatusUnauthorized, "Unauthorized")
        }

        token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("SECRET")), nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, "Unauthorized")
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.JSON(http.StatusUnauthorized, "Unauthorized")
        }

        userID, ok := claims["user_id"].(float64)
        if !ok {
            return c.JSON(http.StatusUnauthorized, "Unauthorized")
        }

        c.Set("user_id", uint(userID))

        return next(c)
    }
}