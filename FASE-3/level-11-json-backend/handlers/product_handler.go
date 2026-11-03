package handlers

import (
	"encoding/json" // Mengolah data json
	"fmt"           // Menampilkan pesan dan formatting
	"io"            // Operasi input/output
	"net/http"      // Membuat HTTP server dan menangani request
	"strings"       // manipulasi string
	"sync"          // Sinkronisasi goroutine

	"level-11-json-backend/models"
)

// Global variabel
var (
	products []models.Product
	nextID   = 1
	mu       sync.Mutex // Mutual Exclusion -> Mengatur agar hanya satu goroutine yang mengakses bagian data yang dilindungi pada satu waktu
)

// Helper untuk mengirim JSON response
// w http.ResponseWriter -> Menulis HTTP response
// status int -> Menentukan HTTP status code
// data any -> Menampung data dengan tipe apa pun
func writeJSON(w http.ResponseWriter, status int, data any) {
	// w -> Response writer.
	// Header() -> Method untuk mengambil HTTP response headers.
	// Set() -> Method untuk menetapkan nilai header.
	w.Header().Set("Content-Type", "application/json")

	// To set the http status code
	w.WriteHeader(status)

	// json.NewEncoder(w) -> Membuat json encode yang menulis hasil ke w
	// .Encode(data) -> Mengubah data mejadi json dan writing to http response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

// GET /products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Lock() -> To restrict (mengunci) akses ke data yang dilindungi
	mu.Lock()
	result := append([]models.Product(nil), products...)

	// To unlock what was previously locked (untuk membukan kembali yang sebelumnya di kunci)
	mu.Unlock()

	writeJSON(w, http.StatusOK, result)
}

// POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	// Membatasi ukuran request body maksimal 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Membuat JSON decoder
	decoder := json.NewDecoder(r.Body)

	// Menolak field JSON yang tidak dikenal
	decoder.DisallowUnknownFields()

	// Mengubah JSON menjadi struct Product
	err := decoder.Decode(&newProduct)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Invalid JSON reqeust body",
		})
		return
	}

	// Memastikan hanya terdapat satu objek JSON
	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Request body must contain only one JSON object",
		})
		return
	}

	// Validasi product name
	newProduct.Name = strings.TrimSpace(newProduct.Name) // Menghapus whitespace diawal dan diakhir

	if newProduct.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product name is required",
		})
		return
	}

	// Validasi price
	if newProduct.Price <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product price must be greater than zero",
		})
		return
	}

	// Validasi stock
	if newProduct.Stock < 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "Product stock cannot be negative",
		})
		return
	}

	// Save product
	mu.Lock()

	newProduct.ID = nextID
	nextID++

	products = append(products, newProduct)

	mu.Unlock()

	// Send response dengan status 201 create
	writeJSON(w, http.StatusCreated, newProduct)
}
