package router

import (
	"user-crud/controller"
	"github.com/gorilla/mux"
)

func NewRouter(userController *controller.UserController) *mux.Router {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/user", userController.FindAll).Methods("GET")
	v1.HandleFunc("/user/{userId}", userController.FindById).Methods("GET")
	v1.HandleFunc("/user", userController.Create).Methods("POST")
	v1.HandleFunc("/user/{userId}", userController.Update).Methods("PATCH")
	v1.HandleFunc("/user/{userId}", userController.Delete).Methods("DELETE")

	return router
}