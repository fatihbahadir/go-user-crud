package main

import (
	"fmt"
	"net/http"
	"user-crud/config"
	"user-crud/controller"
	"user-crud/helper"
	"user-crud/middleware"
	"user-crud/repository"
	"user-crud/router"
	"user-crud/service"
)

func main() {
	fmt.Printf("Server started")

	db := config.DatabaseConnection()

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserServiceImpl(userRepository)

	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	allowedOrigins := []string{
		"http://localhost:3000",
	}

	corsEnabledRoutes := middleware.CORSMiddleware(allowedOrigins)(routes)

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: corsEnabledRoutes,
	}

	err := server.ListenAndServe()
	helper.HandleError(err, "Failed to start server")
}
