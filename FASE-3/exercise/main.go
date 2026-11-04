package main

import (
	"exercise/middlewares"
	"fmt"
	"net/http"
)

func main() {
	mux := setupRoutes()

	handler := middlewares.Logger(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("=======================================")
	fmt.Println("Server running at http://localhost:8080")
	fmt.Println("=======================================")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server Error:", err)
	}
}
