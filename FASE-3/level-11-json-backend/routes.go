package main

import (
	"net/http"

	"level-11-json-backend/handlers"
)

func setupRoutes() *http.ServeMux {
	// Membuat router/multiplexer baru untuk mencocokan URL request dan HTTP method dengan handler-nya.
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /products", handlers.GetProducts)
	mux.HandleFunc("POST /products", handlers.CreateProduct)
	return mux
}
