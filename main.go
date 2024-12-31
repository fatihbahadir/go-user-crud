package main

import (
	"fmt"
	"net/http"
	"user-crud/helper"
)

func main() {
	fmt.Printf("Server started")

	server := http.Server{
		Addr: "localhost:8888",
		Handler: nil,
	}

	err := server.ListenAndServe()
	helper.HandleError(err, "Failed to start server")
}