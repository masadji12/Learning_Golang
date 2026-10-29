// encoding/json		= Mengubah data Go menjadi JSON
// map[string]string 	= Menyimpan pasangan key dan value berupa string
// w.Header().Set()		= Mengatur header response HTTP
// w.WriteHeader() 		= Mengatur status code HTTP
// json.NewEncoder(w).Encode() = Menulis data JSON ke response

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Atur Content-Type menjadi application/json
	w.Header().Set("Content-Type", "application/json") // Memberitahu client bahwa response dari server memiliki format json

	// 2. Atur HTTP status code menjadi 200 OK
	w.WriteHeader(http.StatusOK)

	// 3. Buat map untuk menyimpan response
	response := map[string]string{
		"message": "Product list retrieved successfully",
		"status": "success",
	}

	// 4. Encode data Go menjadi JSON dan kirim ke client
	// json.NewEncoder(w) = Membuat encoder yang menulis ke HTTP response
	// .Encode(response) = Mengubah map menjadi JSON dan mengirimkannya ke client
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func main() {
	// Daftarkan endpoint /products
	http.HandleFunc("/products", productsHandler)

	// Jalankan HTTP server pada port 8080
	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}