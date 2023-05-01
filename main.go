package main

import (
	"go-test/controller"
	"go-test/db"
	"go-test/repository"
	"go-test/usecase"
	"go-test/router"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsercase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsercase)

	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))	
}