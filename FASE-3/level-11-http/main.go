package main

import (
	"fmt"
	"net/http"
)

// http.ResponseWriter = Menulis response kepada client
// *http.Request = Membaca informasi request dari client
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Backend G0!")
}

func main() {
	// http.HandleFunc() = Mendaftarkan handler untuk suatu path
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server berjalan di http://localhost:8080")

	// http.ListenAndServe() = Menjalankan HTTP server pada port tertentu
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}