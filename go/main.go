package main

import (
	"github.com/shiori-42/textbook_change_app/controller"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/router"
	"github.com/shiori-42/textbook_change_app/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.New(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
