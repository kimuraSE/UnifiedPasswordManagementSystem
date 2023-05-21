package main

import (
	"fmt"
	"go-test/controller"
	"go-test/db"
	"go-test/repository"
	"go-test/router"
	"go-test/usecase"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db)
	userUsercase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsercase)

	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
