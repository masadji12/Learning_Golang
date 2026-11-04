package handlers

import (
	"encoding/json"
	"exercise/models"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	products []models.Products
	nextID   = 1
	mu       sync.Mutex
)

func writerJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding json:", err)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	results := append([]models.Products(nil), products...)

	mu.Unlock()

	writerJSON(w, http.StatusOK, results)
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Products

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newProduct)
	if err != nil {
		writerJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Invalid JSON request body",
		})
		return
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		writerJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Request body be must content only one object",
		})
		return
	}

	newProduct.Name = strings.TrimSpace(newProduct.Name)
	if newProduct.Name == "" {
		writerJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product name is required",
		})
		return
	}

	if newProduct.Price <= 0 {
		writerJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product price be must greater than zero",
		})
		return
	}

	if newProduct.Stock < 0 {
		writerJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product stock can't be negative",
		})
		return
	}

	mu.Lock()

	newProduct.ID = nextID
	nextID++

	products = append(products, newProduct)

	mu.Unlock()

	writerJSON(w, http.StatusCreated, newProduct)
}
