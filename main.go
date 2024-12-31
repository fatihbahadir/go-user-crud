package main

import (
	"fmt"
	"net/http"
	"user-crud/config"
	"user-crud/controller"
	"user-crud/helper"
	"user-crud/repository"
	"user-crud/service"
	"user-crud/router"	

)

func main() {
	fmt.Printf("Server started")

	db := config.DatabaseConnection()

	userRepository := repository.NewUserRepository(db)
	
	userService := service.NewUserServiceImpl(userRepository)

	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)	

	server := http.Server{
		Addr: "localhost:8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.HandleError(err, "Failed to start server")
}