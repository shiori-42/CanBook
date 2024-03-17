package main

import (

	"github.com/shiori-42/textbook_change_app/controller"
	"github.com/shiori-42/textbook_change_app/db"
	"github.com/shiori-42/textbook_change_app/repository"
	"github.com/shiori-42/textbook_change_app/router"
	"github.com/shiori-42/textbook_change_app/usecase"
	"github.com/shiori-42/textbook_change_app/validator"
			
)

func main() {
	db := db.NewDB()
	userValidator:=validator.NewUserValidator()
	listingValidator:=validator.NewListingValidator()
	userRepository := repository.NewUserRepository(db)
	listingRepository := repository.NewListingRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository,userValidator)
	listingUsecase := usecase.NewListingUsecase(listingRepository,listingValidator)
	userController := controller.NewUserController(userUsecase)
	listingController := controller.NewListingController(listingUsecase)
	e := router.NewRouter(userController,listingController)
	e.Logger.Fatal(e.Start(":8080"))
}
