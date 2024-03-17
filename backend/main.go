package main

import (
	"container/list"

	"github.com/shiori-42/textbook_change_app/controller"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/router"
	"github.com/shiori-42/textbook_change_app/usecase"
)

func main() {
	db := db.NewDB()
	userValidator:=validator.NewUserValidator()
	listingValidator:=validator.NewListingValidator()
	userRepository := repository.NewUserRepository(db)
	lisgingRepository := repository.NewListingRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository,userValidator)
	listingUsecase := usecase.NewListingUsecase(lisgingRepository,lisgingvalidator)
	userController := controller.NewUserController(userUsecase)
	listingController := controller.NewListingController(listingUsecase)
	e := router.NewRouter(userController,listingController)
	e.Logger.Fatal(e.Start(":8080"))
}
