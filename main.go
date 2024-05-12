package main

import (
	"todo-deck-api/controller"
	"todo-deck-api/db"
	"todo-deck-api/repository"
	"todo-deck-api/routes"
	"todo-deck-api/usecase"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userUsecase)

	e := routes.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}