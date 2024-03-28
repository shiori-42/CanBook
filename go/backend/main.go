/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/22 08:18:06 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/28 02:42:31 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/shiori-42/textbook_change_app/go/backend/api"
	"github.com/shiori-42/textbook_change_app/go/backend/db"
	"github.com/shiori-42/textbook_change_app/go/backend/model"
)

func main() {
	if os.Getenv("GO_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
			return
		}
	}

	if err := db.InitDB(); err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.CloseDB()

	if err := model.AutoMigrate(); err != nil {
		fmt.Printf("Failed to auto migrate: %v\n", err)
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	frontURL := os.Getenv("FE_URL")
	if frontURL == "" {
		frontURL = "http://localhost:3000"
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderXCSRFToken, echo.HeaderAuthorization, echo.HeaderAccessControlAllowHeaders,
			"Upgrade", "Connection"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	api.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
