package main

import (
	"fmt"
)

func main() {
	categories := [3]string {
		"Alat Tulis",
		"Kertas",
		"Perlengkap IT",
	}

	products := []string {
		"Pulpen",
		"Buku",
		"Pensil",
	}

	stocks := map[string]int {
		"Pulpen": 20,
		"Buku": 15,
		"Pensil": 30,
	}

	fmt.Println("\n=== CATEGORY ATK ===")

	for i, category := range categories {
		fmt.Println(i+1, category)
	}

	fmt.Println("\n=== DAFTAR INVENTARIS ===")

	for i, product := range products {
		fmt.Println("No:", i, "| Barang", product, "| Stock", stocks[product])
	}

	totalstock := 0

	for _, product := range products {
		totalstock += stocks[product]
	}

	fmt.Println("\n=== RINGKASAN ===")
	fmt.Println("Jumlah jenis producT:", len(products))
	fmt.Println("Total unit stok:", totalstock)
}