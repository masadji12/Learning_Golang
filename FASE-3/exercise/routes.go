package main

import (
	"exercise/handlers"
	"net/http"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", handlers.GetProducts)
	mux.HandleFunc("POST /products", handlers.CreateProducts)

	return mux
}
