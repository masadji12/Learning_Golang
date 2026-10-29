package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID		int 	`json:"id"`
	Name	string	`json:"name"`
	Price	float64	`json:"price"`
	Stock	int		`json:"stock"`
}

// Temporary storage in memory (Penyimpan sementara di memori)
var products []Product
var nextID = 1

func productsHandler(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// GET: mengambil list product
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(products); err != nil {
			fmt.Println("Error encoding products:", err)
		}
		
	case http.MethodPost:
		// POST: menerima JSON dari request body
		var newProduct Product

		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(
				w,
				`{"message":"Invalid JSON request body"}`,
				http.StatusBadRequest,
			)
			return
		}

		// Validasi sederhana
		if newProduct.Name == "" || newProduct.Price < 0 || newProduct.Stock < 0 {
			http.Error(
				w,
				`{"message":"Invalid product data}`,
				http.StatusBadRequest,
			)
			return
		}

		// Server menentukan ID product
		newProduct.ID = nextID
		nextID++

		// Simpan product ke slice
		products = append(products, newProduct)

		// Kirim product yang berhasil dibuat
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(newProduct); err != nil {
			fmt.Println("Error encoding new product:", err)
		}

	default:
		// Method selain GET dan POST tidak didukung
		w.Header().Set("Allow", "GET, POST")
		http.Error(
			w,
			`{"message":"Method not allowed"}`,
			http.StatusMethodNotAllowed,
		)
	}
}

func main () {
	http.HandleFunc("/products", productsHandler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}