package main

import (
	"fmt"
	"log"
	"net/http"

	"level-11-json-backend/middleware"
)

func main() {
	// Mendaftarkan Seluruh endpoint
	mux := setupRoutes()

	// Mengaktifkan middleware logger
	handler := middleware.Logger(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("================================")
	fmt.Println(" Go HTTP API Server")
	fmt.Println(" Server running at :8080")
	fmt.Println("================================")

	// Running server
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server error:", err)
	}
}
