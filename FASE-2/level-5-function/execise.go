package main

import "fmt"

func tampilkanBarang(products []string, stocks map[string]int) {
	fmt.Println("=== LIST PRODUCTS ===")

	for _, product := range products {
		fmt.Println("Barang:", product, "| Stock:", stocks[product])
	}
}

func hitungTotalStock(products []string, stocks map[string]int) int {
	total := 0

	for _, product := range products {
		total += stocks[product]
	}

	return total
}

func main() {
	products := []string{"Pulpen", "Kertas", "Penggaris"}

	stocks := map[string]int{
		"Pulpen": 20,
		"Kertas": 30,
		"Penggaris": 50,
	}

	tampilkanBarang(products, stocks)

	total := hitungTotalStock(products, stocks)
	fmt.Println("\nTotal Stocks:", total)
}