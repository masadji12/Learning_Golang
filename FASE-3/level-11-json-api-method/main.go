package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductsPatch struct {
	ID		*int	`json:"id"`
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	Stock *int     `json:"stock"`
}

var (
	products []Product
	nextID   = 1
	mux      sync.RWMutex
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"message": message,
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s - %v",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}

func findProductIndex(id int) int {
	for i, product := range products {
		if product.ID == id {
			return i
		}
	}
	return -1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mux.RLock()
		results := append([]Product(nil), products...)
		mux.RUnlock()
		writeJSON(w, http.StatusOK, results)

	case http.MethodPost:
		var newProduct Product

		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON request body")
			return
		}

		newProduct.Name = strings.TrimSpace(newProduct.Name)
		if newProduct.Name == "" || newProduct.Price <= 0 || newProduct.Stock < 0 {
			writeError(w, http.StatusBadRequest, "Invalid Product data")
			return
		}

		mux.Lock()
		newProduct.ID = nextID
		nextID++

		products = append(products, newProduct)
		mux.Unlock()

		writeJSON(w, http.StatusCreated, newProduct)

	default:
		w.Header().Set("Allow", "GET, POST")
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func productHandlerByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		mux.RLock()
		index := findProductIndex(id)
		if index == -1 {
			mux.RUnlock()
			writeError(w, http.StatusBadRequest, "Product not found")
			return
		}
		product := products[index]
		mux.RUnlock()
		
		writeJSON(w, http.StatusOK, product)

	case http.MethodPut:
		var updatedProduct Product

		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON request body")
			return
		}

		updatedProduct.Name = strings.TrimSpace(updatedProduct.Name)
		if updatedProduct.Name == "" || updatedProduct.Price <= 0 || updatedProduct.Stock < 0 {
			writeError(w, http.StatusBadRequest, "Invalid Product data")
			return
		}

		mux.Lock()
		index := findProductIndex(id)

		if index == -1 {
			mux.Unlock()
			writeError(w, http.StatusBadRequest, "Product not found")
			return
		}

		updatedProduct.ID = id
		products[index] = updatedProduct
		mux.Unlock()

		writeJSON(w, http.StatusOK, updatedProduct)

	case http.MethodPatch:
		var patch ProductsPatch

		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON request body")
			return
		}

		// Validation
		if patch.Name != nil {
			*patch.Name = strings.TrimSpace(*patch.Name)
			if *patch.Name == "" {
				writeError(w, http.StatusBadRequest, "Product name can't be empty")
				return
			}
		}

		if patch.Price != nil && *patch.Price <= 0 {
			writeError(w, http.StatusBadRequest, "Product price must be greate than zero")
			return
		}

		if patch.Stock != nil && *patch.Stock < 0 {
			writeError(w, http.StatusBadRequest, "Product stock can't be negative")
			return
		}

		mux.Lock()
		index := findProductIndex(id)

		if index == -1 {
			mux.Unlock()
			writeError(w, http.StatusBadRequest, "Product not found")
			return
		}

		if patch.Name != nil {
			products[index].Name = *patch.Name
		}

		if patch.Price != nil {
			products[index].Price = *patch.Price
		}

		if patch.Stock != nil {
			products[index].Stock = *patch.Stock
		}

		updatedProduct := products[index]
		mux.Unlock()

		writeJSON(w, http.StatusOK, updatedProduct)

	case http.MethodDelete:
		mux.Lock()
		index := findProductIndex(id)
		if index == -1 {
			mux.Unlock()
			writeError(w, http.StatusBadRequest, "Product not found")
			return
		}

		products = append(products[:index], products[index+1:]...)
		mux.Unlock()

		writeJSON(w, http.StatusOK, map[string]string {
			"message": "Product deleted successfully",
		})

	default:
		w.Header().Set("Allow", "GET, PUT, PATCH, DELETE")
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", productsHandler)
	mux.HandleFunc("/products/{id}", productHandlerByID)

	handler := Logger(mux)
	
	fmt.Println("Server running at http://localhost:8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error", err)
	}
}
