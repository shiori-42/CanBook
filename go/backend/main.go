/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: shiori0123 <shiori0123@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/03/20 12:04:16 by shiori0123        #+#    #+#             */
/*   Updated: 2024/03/20 14:26:07 by shiori0123       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/shiori-42/textbook_change_app/api"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/model"
	"os"
)

func main() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	if err := model.AutoMigrate(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully Migrated")
	defer db.CloseDB()
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)

	api.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
